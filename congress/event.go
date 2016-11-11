package congress

import (
  "log"

  "github.com/30x/congress/policy"
  "github.com/30x/congress/utils"

  "k8s.io/kubernetes/pkg/api"
  "k8s.io/kubernetes/pkg/client/unversioned"
  "k8s.io/kubernetes/pkg/watch"
  ext "k8s.io/kubernetes/pkg/apis/extensions"
)

func HandleNamespaceEvent(event watch.Event, kubeClient *unversioned.Client, extClient *unversioned.ExtensionsClient, config *utils.Config) (err error) {
  namespace := event.Object.(*api.Namespace)

  if event.Type == watch.Added {
    // determine if congress is ignoring/excluding this namespace
    if !config.IsExcluded(namespace.Name) && !config.InIgnoreSelector(namespace.GetLabels()) {
      log.Printf("New namespace added: %s\n", namespace.Name)

      // add label for ingress policy ID
      if config.IsolateNamespace {
        err := policy.IsolateNamespace(kubeClient, namespace, config)
        if err != nil { return err }
      }

      // add network policies
      err = policy.EnactPolicies(extClient, namespace, config)
      if err != nil { return err }
    }
  } else if event.Type == watch.Modified {
    // determine if congress is ignoring/excluding this namespace
    if namespace.Status.Phase == api.NamespaceActive && !config.IsExcluded(namespace.Name) && !config.InIgnoreSelector(namespace.GetLabels()) {
      log.Printf("Modification on namespace: %s. Validating", namespace.Name)

      // verify any namespace modifications didn't remove our policies
      err = policy.ValidateNamespace(kubeClient, extClient, namespace, config)
      if err != nil { return err }
    } else if config.InIgnoreSelector(namespace.GetLabels()) && namespace.Status.Phase == api.NamespaceActive {
      log.Printf("Modification on excluded namespace: %s. Ensuring it isn't isolated.", namespace.Name)

      // this namespace might be newly excluded, ensure it is exposed
      err = policy.ValidateIsolation(kubeClient, extClient, namespace, config)
      if err != nil { return err }
    }
  }

  return
}

// HandlePolicyEvent handles events on network policy objects
func HandlePolicyEvent(event watch.Event, kubeClient *unversioned.Client, extClient *unversioned.ExtensionsClient, config *utils.Config) (err error) {
  policyObj := event.Object.(*ext.NetworkPolicy)

  if event.Type == watch.Deleted {
    namespace, err := kubeClient.Namespaces().Get(policyObj.ObjectMeta.Namespace)
    if err != nil { return err }

    // determine if congress is ignoring/excluding this namespace
    if !config.IsExcluded(namespace.Name) && !config.InIgnoreSelector(namespace.GetLabels()) && namespace.Status.Phase != api.NamespaceTerminating && policy.IsCongressPolicy(policyObj.Name, config) {
      log.Println("Attempt to delete policy that shouldn't be deleted! Reenacting...")

      // reenact the policy (if it is one of ours)
      err = policy.ReenactPolicy(policyObj, namespace, extClient, config)
      if err != nil { return err }

      log.Printf("Reenacted policy \"%s\" in namespace \"%s\"\n", policyObj.ObjectMeta.Name, namespace.Name)
    }
  }

  return
}