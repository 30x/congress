package congress

import (
  "log"

  "github.com/30x/congress/policy"
  "github.com/30x/congress/utils"

  "k8s.io/kubernetes/pkg/api"
  "k8s.io/kubernetes/pkg/client/unversioned"
  "k8s.io/kubernetes/pkg/watch"
)

// InitCongress starts a namespace watcher
func InitCongress(client *unversioned.Client, extClient *unversioned.ExtensionsClient, config *utils.Config) (namespaceWatcher watch.Interface, policyWatcher watch.Interface) {
  log.Println("Retrieving watchers")

  namespaceWatcher = StartNamespaceWatcher(client, extClient, config)

  policyWatcher = StartPolicyWatcher(extClient, config)

  return
}

func StartNamespaceWatcher(client *unversioned.Client, extClient *unversioned.ExtensionsClient, config *utils.Config) (namespaceWatcher watch.Interface) {
  namespaces := client.Namespaces()

  existingNamespaces, err := namespaces.List(api.ListOptions{
    LabelSelector: config.LabelSelector,
  })
  if err != nil {
    log.Fatalf("Could not get list of existing namespaces")
  }

  log.Println("Validating existing namespaces")
  validated, err := policy.ValidateList(client, extClient, existingNamespaces, config)
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

  namespaceWatcher, err = namespaces.Watch(nsWatchOpts)
  if err != nil {
    log.Fatalf("Failed making namespace watcher: %v", err)
  }

  return
}

func StartPolicyWatcher(extClient *unversioned.ExtensionsClient, config *utils.Config) (policyWatcher watch.Interface) {
  policies := extClient.NetworkPolicies(api.NamespaceAll)

  existingPolicies, err := policies.List(api.ListOptions{})
  if err != nil {
    log.Fatalf("Couldn't get list of existing network policies: %v", err)
  }

  policyWatchOpts := api.ListOptions{
    ResourceVersion: existingPolicies.ResourceVersion,
  }

  policyWatcher, err = policies.Watch(policyWatchOpts)
  if err != nil {
    log.Fatalf("Failed making policy watcher: %v", err)
  }

  return
}