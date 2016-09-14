// Copyright © 2016 Apigee Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
  "log"

  "github.com/30x/congress/utils"
  "github.com/30x/congress/policy"

  "k8s.io/kubernetes/pkg/api"
  "k8s.io/kubernetes/pkg/client/unversioned"
  "k8s.io/kubernetes/pkg/watch"
)

func main() {
  var doRestart bool

  // get k8s client
  kubeClient, err := utils.GetClient()
  if err != nil {
    log.Fatalf("Failed to create client: %v", err)
  }

  // get extensions client
  extClient, err := utils.GetExtensionsClient()
  if err != nil {
    log.Fatalf("Failed to create extensions client: %v", err)
  }

  // get congress configuration from environment
  config, err := utils.ConfigFromEnv()
  if err != nil {
    log.Fatalf("Failed to retrieve the congress config: %v", err)
  }

  // start a namespace watcher on the cluster
  watcher := initCongress(kubeClient, extClient, config)
  log.Println("Congress is in session")

  // begin reading namespace events
  eventCh := watcher.ResultChan()
  for {
    // block until namespace event occurs
    event, ok := <- eventCh
    if !ok {
      log.Println("Kuberenetes closed the namespace watcher")
      doRestart = true
    }

    // new namespace created event, start isolating it
    if event.Type == watch.Added && !doRestart {
      namespace := event.Object.(*api.Namespace)
      if !config.IsExcluded(namespace.Name) && !config.InIgnoreSelector(namespace.GetLabels()) {
        log.Printf("New namespace added: %s\n", namespace.Name)

        // add label for ingress policy ID
        if config.IsolateNamespace {
          err := policy.IsolateNamespace(kubeClient, namespace)
          if err != nil {
            log.Printf("Failed writing namespace ingress isolation: %v", err)
          }
        }

        err = policy.EnactPolicies(extClient, namespace, config)
        if err != nil {
          log.Printf("Failed enacting network policies: %v", err)
        }
      }
    } else if event.Type == watch.Modified && !doRestart {
      // verify any namespace modifications didn't remove our policies
      namespace := event.Object.(*api.Namespace)
      if namespace.Status.Phase == api.NamespaceActive && !config.IsExcluded(namespace.Name) &&
        !config.InIgnoreSelector(namespace.GetLabels()) {
        log.Printf("Modification on namespace: %s. Validating", namespace.Name)
        err = policy.ValidateNamespace(kubeClient, extClient, namespace, config)
        if err != nil {
          log.Printf("Failed validating modified namespace %s: %v", namespace.Name, err)
        }
      }
    }

    if doRestart {
      log.Println("Congress is restarting the watcher")
      doRestart = false
      watcher.Stop()
      watcher = initCongress(kubeClient, extClient, config)
      eventCh = watcher.ResultChan()
    }
  }
}

// initCongress starts a namespace watcher
func initCongress(client *unversioned.Client, extClient *unversioned.ExtensionsClient, config *utils.Config) (watch.Interface) {
  log.Println("Retrieving watcher")

  namespaces := client.Namespaces()

  existing, err := namespaces.List(api.ListOptions{
    LabelSelector: config.LabelSelector,
  })
  if err != nil {
    log.Printf("Could not get list of existing namespaces")
  }

  log.Println("Validating existing namespaces")
  validated, err := policy.ValidateList(client, extClient, existing, config)
  if err != nil {
    log.Fatalf("Failed validating the existing Namespaces list: %v", err)
  }

  if config.LabelSelector.String() != "" {
    log.Printf("Selecting namespaces on: %v", config.LabelSelector)
  }

  if config.IgnoreSelector.String() != "" {
    log.Printf("Ignoring namespaces on: %v", config.IgnoreSelector)
  }

  nsWatchOpts := api.ListOptions{
    LabelSelector: config.LabelSelector,
    ResourceVersion: validated.ResourceVersion,
  }

  watcher, err := namespaces.Watch(nsWatchOpts)
  if err != nil {
    log.Fatalf("Failed making namespace watcher: %v", err)
  }

  return watcher
}

