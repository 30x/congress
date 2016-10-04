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
    if !config.IsExcluded(ns.Name) && !config.InIgnoreSelector(ns.GetLabels()) {
      err := ValidateNamespace(client, extClient, &ns, config)
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
func ValidateNamespace(client *unversioned.Client, extClient *unversioned.ExtensionsClient, namespace *api.Namespace, config *utils.Config) error {
  updated, err := validateAnnotationAndLabel(client, namespace, config)
  if err != nil {
    return err
  }

  err = validateNetworkPolicies(extClient, updated, config)
  if err != nil {
    return err
  }

  return nil
}

// ValidateIsolation ensures that a namespace excluded by the IgnoreSelector labels is not isolated by congress
func ValidateIsolation(client *unversioned.Client, extClient *unversioned.ExtensionsClient, namespace *api.Namespace, config *utils.Config) error {
  updated, err := repealIsolationAnnotationAndLabel(client, namespace, config)
  if err != nil {
    return err
  }

  return repealNetworkPolicies(extClient, updated, config)
}

func validateNetworkPolicies(extClient *unversioned.ExtensionsClient, namespace *api.Namespace, config *utils.Config) error {
  policies := extClient.NetworkPolicies(namespace.Name)

  _, err := policies.Get(IntraPolicyName)
  if err != nil { // didn't have the allow-intra-namespace policy
    log.Printf("%s missing %s policy. Adding it to validate.", namespace.Name, IntraPolicyName)
    err = AddIntraPolicy(extClient, namespace, config)
    if err != nil {
      return err
    }
  }

  _, err = policies.Get(config.RoutingPolicyName)
  if err != nil { // didn't have the allow-apigee policy
    log.Printf("%s missing %s policy. Adding it to validate.", namespace.Name, config.RoutingPolicyName)
    err = AddBridgePolicy(extClient, namespace, config)
    if err != nil {
      return err
    }
  }

  return nil
}

func validateAnnotationAndLabel(client *unversioned.Client, namespace *api.Namespace, config *utils.Config) (*api.Namespace, error) {
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

  if val, exists := labels[config.RoutingLabelName]; !exists ||
    (exists && val != namespace.Name) {

    // add `name` label with this namespace's name
    labels[config.RoutingLabelName] = namespace.Name
    namespace.SetLabels(labels)
    doUpdate = true
  }

  if doUpdate {
    updated, err := client.Namespaces().Update(namespace)
    if err != nil {
      return nil, err
    }

    // reset flag, just in case
    doUpdate = false

    log.Printf("%s was invalid. Updated to validate.", namespace.Name)
    return updated, nil
  }

  return namespace, nil
}

func repealIsolationAnnotationAndLabel(client *unversioned.Client, namespace *api.Namespace, config *utils.Config) (*api.Namespace, error) {
  var doUpdate bool

  annotations := namespace.GetAnnotations()
  if annotations == nil { // no annotations, don't need to do anything
    return namespace, nil
  }

  if annotations != nil {
    if val, exists := annotations[IngressAnnotationKey]; exists && val == IngressAnnotationValue {

      // annotation does exist, remove it
      delete(annotations, IngressAnnotationKey)

      namespace.SetAnnotations(annotations)
      doUpdate = true
    }
  }

  labels := namespace.GetLabels()

  if labels != nil {
    if val, exists := labels[config.RoutingLabelName]; exists && val == namespace.Name {

      // repeal routing label from namespace
      delete(labels, config.RoutingLabelName)

      namespace.SetLabels(labels)
      doUpdate = true
    }
  }

  if doUpdate {
    updated, err := client.Namespaces().Update(namespace)
    if err != nil {
      return nil, err
    }

    // reset flag, just in case
    doUpdate = false

    log.Printf("%s was isolated while excluded. Removing isolation properties.", namespace.Name)
    return updated, nil
  }

  return namespace, nil
}

func repealNetworkPolicies(extClient *unversioned.ExtensionsClient, namespace *api.Namespace, config *utils.Config) error {
  policies := extClient.NetworkPolicies(namespace.Name)

  _, err := policies.Get(IntraPolicyName)
  if err == nil { // has the allow-intra-namespace policy, repeal it
    log.Printf("%s has %s policy. Repealing it.", namespace.Name, IntraPolicyName)
    err = policies.Delete(IntraPolicyName, nil)
    if err != nil {
      return err
    }
  }

  _, err = policies.Get(config.RoutingPolicyName)
  if err == nil { // has the allow-apigee policy, repeal it
    log.Printf("%s has %s policy. Repealing it.", namespace.Name, config.RoutingPolicyName)
    err = policies.Delete(config.RoutingPolicyName, nil)
    if err != nil {
      return err
    }
  }

  return nil
}