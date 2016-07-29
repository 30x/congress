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

package policy

import (
  "log"

  "github.com/30x/congress/utils"

  "k8s.io/kubernetes/pkg/api"
  "k8s.io/kubernetes/pkg/client/unversioned"
)

// ValidateList validates that a list of namespaces conform to network isolation standards
func ValidateList(client *unversioned.Client, extClient *unversioned.ExtensionsClient, list *api.NamespaceList, config *utils.Config) (*api.NamespaceList, error) {
  for _, ns := range list.Items {
    if !config.IsExcluded(ns.Name) {
      err := ValidateNamespace(client, extClient, &ns)
      if err != nil {
        return nil, err
      }
    }
  }

  list, err := client.Namespaces().List(api.ListOptions{
    LabelSelector: config.LabelSelector,
  })
  if err != nil {
    return nil, err
  }

  return list, nil
}

// ValidateNamespace validates that a single namespace conforms to network isolation standards
func ValidateNamespace(client *unversioned.Client, extClient *unversioned.ExtensionsClient, namespace *api.Namespace) error {
  updated, err := validateAnnotationAndLabel(client, namespace)
  if err != nil {
    return err
  }

  err = validateNetworkPolicies(extClient, updated)
  if err != nil {
    return err
  }

  return nil
}

func validateNetworkPolicies(extClient *unversioned.ExtensionsClient, namespace *api.Namespace) error {
  policies := extClient.NetworkPolicies(namespace.Name)

  _, err := policies.Get(IntraPolicyName)
  if err != nil { // didn't have the allow-intra-namespace policy
    log.Printf("%s missing %s policy. Adding it to validate.", namespace.Name, IntraPolicyName)
    err = AddIntraPolicy(extClient, namespace)
    if err != nil {
      return err
    }
  }

  _, err = policies.Get(BridgePolicyName)
  if err != nil { // didn't have the allow-apigee policy
    log.Printf("%s missing %s policy. Adding it to validate.", namespace.Name, BridgePolicyName)
    err = AddBridgePolicy(extClient, namespace)
    if err != nil {
      return err
    }
  }

  return nil
}

func validateAnnotationAndLabel(client *unversioned.Client, namespace *api.Namespace) (*api.Namespace, error) {
  var doUpdate bool

  annotations := namespace.GetAnnotations()
  if annotations == nil { // no annotations
    annotations = map[string]string{}
  }

  if val, exists := annotations[IngressAnnotationKey]; !exists ||
    (exists && val != IngressAnnotationValue) {

    // annotation does not exist, add it
    annotations[IngressAnnotationKey] = IngressAnnotationValue
    namespace.SetAnnotations(annotations)
    doUpdate = true
  }

  labels := namespace.GetLabels()
  if labels == nil { // no labels
    labels = map[string]string{}
  }

  if val, exists := labels[NameLabelKey]; !exists ||
    (exists && val != namespace.Name) {

    // add `name` label with this namespace's name
    labels[NameLabelKey] = namespace.Name
    namespace.SetLabels(labels)
    doUpdate = true
  }

  if doUpdate {
    updated, err := client.Namespaces().Update(namespace)
    if err != nil {
      return nil, err
    }

    log.Printf("%s was invalid. Updated to validate.", namespace.Name)
    return updated, nil
  }

  return namespace, nil
}