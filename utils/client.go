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

package utils

import (
  "k8s.io/kubernetes/pkg/client/unversioned"
  clientcmd "k8s.io/kubernetes/pkg/client/unversioned/clientcmd"
)

// GetClient retrieves a k8s client based on the current kube config
func GetClient() (*unversioned.Client, error) {
	// retrieve necessary kube config settings
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	// make a client config with kube config
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	// make a client out of the kube client config
	client, err := unversioned.New(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetExtensionsClient retrieves a k8s extensions client based on current kube config
func GetExtensionsClient() (*unversioned.ExtensionsClient, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	// make a client config with kube config
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	client, err := unversioned.NewExtensions(config)
	if err != nil {
		return nil, err
	}

	return client, nil
}