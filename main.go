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
      if !config.IsExcluded(namespace.Name) {
        log.Printf("New namespace added: %s\n", namespace.Name)

        // add ingress isolation annotation & label for policies to namespace
        err := policy.IsolateNamespace(kubeClient, namespace)
        if err != nil {
          log.Printf("Failed writing ingress annotation: %v", err)
          continue
        }

        err = policy.EnactPolicies(extClient, namespace)
        if err != nil {
          log.Printf("Failed enacting network policies: %v", err)
          continue
        }
      }
    } else if event.Type == watch.Modified && !doRestart {
      // verify any namespace modifications didn't remove our policies
      namespace := event.Object.(*api.Namespace)
      if namespace.Status.Phase == api.NamespaceActive && !config.IsExcluded(namespace.Name) {
        log.Printf("Modification on namespace: %s. Validating", namespace.Name)
        err = policy.ValidateNamespace(kubeClient, extClient, namespace)
        if err != nil {
          log.Printf("Failed validating modified namespace %s", namespace.Name)
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

  existing, err := namespaces.List(api.ListOptions{})
  if err != nil {
    log.Printf("Could not get list of existing namespaces")
  }

  log.Println("Validating existing namespaces")
  validated, err := policy.ValidateList(client, extClient, existing, config)
  if err != nil {
    log.Fatalf("Failed validating the existing Namespaces list: %v", err)
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

