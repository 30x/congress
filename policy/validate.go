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

  list, err := client.Namespaces().List(api.ListOptions{})
  if err != nil {
    return nil, err
  }

  return list, nil
}

// ValidateNamespace validates that a single namespace conforms to network isolation standards
func ValidateNamespace(client *unversioned.Client, extClient *unversioned.ExtensionsClient, namespace *api.Namespace) error {
  err := validateAnnotationAndLabel(client, namespace)
  if err != nil {
    return err
  }

  err = validateNetworkPolicies(extClient, namespace)
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

func validateAnnotationAndLabel(client *unversioned.Client, namespace *api.Namespace) error {
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
    _, err := client.Namespaces().Update(namespace)
    if err != nil {
      return err
    }

    log.Printf("%s was invalid. Updated to validate.", namespace.Name)
  }

  return nil
}