/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package app implements a server that runs a set of active
// components.  This includes cluster controller

package app

import (
	"net"
	"net/http"
	"net/http/pprof"
	"strconv"

	federationclientset "k8s.io/kubernetes/federation/client/clientset_generated/federation_release_1_3"
	"k8s.io/kubernetes/federation/cmd/federation-controller-manager/app/options"
	"k8s.io/kubernetes/federation/pkg/dnsprovider"
	clustercontroller "k8s.io/kubernetes/federation/pkg/federation-controller/cluster"
	servicecontroller "k8s.io/kubernetes/federation/pkg/federation-controller/service"
	"k8s.io/kubernetes/federation/pkg/federation-controller/util"
	"k8s.io/kubernetes/pkg/client/restclient"
	"k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
	"k8s.io/kubernetes/pkg/healthz"
	"k8s.io/kubernetes/pkg/util/configz"
	"k8s.io/kubernetes/pkg/util/wait"
	"k8s.io/kubernetes/pkg/version"

	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	// "federation-apiserver-secret" is a reserved secret name which stores the kubeconfig for federation-apiserver.
	FederationAPIServerSecretName = "federation-apiserver-secret"
)

// NewControllerManagerCommand creates a *cobra.Command object with default parameters
func NewControllerManagerCommand() *cobra.Command {
	s := options.NewCMServer()
	s.AddFlags(pflag.CommandLine)
	cmd := &cobra.Command{
		Use: "federation-controller-manager",
		Long: `The federation controller manager is a daemon that embeds
the core control loops shipped with federation. In applications of robotics and
automation, a control loop is a non-terminating loop that regulates the state of
the system. In federation, a controller is a control loop that watches the shared
state of the federation cluster through the apiserver and makes changes attempting
to move the current state towards the desired state. Examples of controllers that
ship with federation today is the cluster controller.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}

// Run runs the CMServer.  This should never exit.
func Run(s *options.CMServer) error {
	glog.Infof("%+v", version.Get())
	if c, err := configz.New("componentconfig"); err == nil {
		c.Set(s.ControllerManagerConfiguration)
	} else {
		glog.Errorf("unable to register configz: %s", err)
	}
	// Create the config to talk to federation-apiserver.
	kubeconfigGetter := util.KubeconfigGetterForSecret(FederationAPIServerSecretName)
	restClientCfg, err := clientcmd.BuildConfigFromKubeconfigGetter(s.Master, kubeconfigGetter)
	if err != nil {
		return err
	}

	// Override restClientCfg qps/burst settings from flags
	restClientCfg.QPS = s.APIServerQPS
	restClientCfg.Burst = s.APIServerBurst

	go func() {
		mux := http.NewServeMux()
		healthz.InstallHandler(mux)
		if s.EnableProfiling {
			mux.HandleFunc("/debug/pprof/", pprof.Index)
			mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
			mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		}
		mux.Handle("/metrics", prometheus.Handler())

		server := &http.Server{
			Addr:    net.JoinHostPort(s.Address, strconv.Itoa(s.Port)),
			Handler: mux,
		}
		glog.Fatal(server.ListenAndServe())
	}()

	run := func() {
		err := StartControllers(s, restClientCfg)
		glog.Fatalf("error running controllers: %v", err)
		panic("unreachable")
	}
	run()
	panic("unreachable")
}

func StartControllers(s *options.CMServer, restClientCfg *restclient.Config) error {

	ccClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, "cluster-controller"))
	go clustercontroller.NewclusterController(ccClientset, s.ClusterMonitorPeriod.Duration).Run()
	dns, err := dnsprovider.InitDnsProvider(s.DnsProvider, s.DnsConfigFile)
	if err != nil {
		glog.Fatalf("Cloud provider could not be initialized: %v", err)
	}
	scClientset := federationclientset.NewForConfigOrDie(restclient.AddUserAgent(restClientCfg, servicecontroller.UserAgentName))
	servicecontroller := servicecontroller.New(scClientset, dns, s.FederationName, s.ZoneName)
	if err := servicecontroller.Run(s.ConcurrentServiceSyncs, wait.NeverStop); err != nil {
		glog.Errorf("Failed to start service controller: %v", err)
	}
	select {}
}
