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
  "k8s.io/kubernetes/pkg/api"
  unversionedApi "k8s.io/kubernetes/pkg/api/unversioned"
  "k8s.io/kubernetes/pkg/client/unversioned"
  "k8s.io/kubernetes/pkg/apis/extensions"

  "github.com/30x/congress/utils"
)

const (
  // IngressAnnotationKey key for the network policy annotation in a namespace
  IngressAnnotationKey = "net.beta.kubernetes.io/network-policy"
  // IngressAnnotationValue is the policy that belongs to the NetworkPolicy key
  IngressAnnotationValue = `{"ingress": {"isolation": "DefaultDeny"}}"`
  // IntraPolicyName name of the intra-namespace network policy
  IntraPolicyName = "allow-intra-namespace"
  // BridgePolicyName name of network policy that allows apigee pod traffic to the namespace
  BridgePolicyName = "allow-apigee"
)

// IsolateNamespace adds the necessary label for network isolation
func IsolateNamespace(client *unversioned.Client, namespace *api.Namespace, config *utils.Config) error {
  addIngressAnnotation(namespace)
  addRoutingLabel(namespace, config.RoutingLabelName)

  _ , err := client.Namespaces().Update(namespace)
  if err != nil {
    return err
  }

  return nil
}

func addIngressAnnotation(namespace *api.Namespace) {
  annotations := namespace.GetAnnotations()
  if annotations == nil { // no annotations
    annotations = map[string]string{}
  } else if _, exists := annotations[IngressAnnotationKey]; exists {
    return // annotation already exists
  }

  // annotation does not exist, add it
  annotations[IngressAnnotationKey] = IngressAnnotationValue
  namespace.SetAnnotations(annotations)

  return
}

func addRoutingLabel(namespace *api.Namespace, routingLabelName string) {
  labels := namespace.GetLabels()
  if labels == nil { // no labels
    labels = map[string]string{}
  } else if _, exists := labels[routingLabelName]; exists {
    return // label already exists
  }

  // add `name` label with this namespace's name
  labels[routingLabelName] = namespace.Name
  namespace.SetLabels(labels)
}

// EnactPolicies creates the necessary network policies in the given namespace
func EnactPolicies(client *unversioned.ExtensionsClient, namespace *api.Namespace, config *utils.Config) error {
  policies := client.NetworkPolicies(namespace.Name)

  _, err := policies.Get(IntraPolicyName)
  if err != nil { // policy did not exist in namespace, make it
    err = AddIntraPolicy(client, namespace, config)
    if err != nil {
      return err
    }
  } else {
    log.Printf("%s already exists in %s. Skipping creation.", IntraPolicyName, namespace.Name)
  }

  _, err = policies.Get(BridgePolicyName)
  if err != nil { // policy did not exist in namespace, make it
    err = AddBridgePolicy(client, namespace, config)
    if err != nil {
      return err
    }
  } else {
    log.Printf("%s already exists in %s. Skipping creation.", BridgePolicyName, namespace.Name)
  }

  return nil
}

// AddBridgePolicy adds the allow-apigee NetworkPolicy to the given namespace
func AddBridgePolicy(client *unversioned.ExtensionsClient, namespace *api.Namespace, config *utils.Config) error {
  // get the apigee bridge network policy
  bridge := writeBridgePolicy(config)

  // add the allow-apigee NetworkPolicy to the namespace
  _, err :=  client.NetworkPolicies(namespace.Name).Create(bridge)
  if err != nil {
    return err
  }

  return nil
}

// AddIntraPolicy adds the intra-namespace network policy to a given namespace
func AddIntraPolicy(client *unversioned.ExtensionsClient, namespace *api.Namespace, config *utils.Config) error {
  // make intra-namespace policy for this namespace
  intra := writeIntraPolicy(namespace, config.RoutingLabelName)

  // add allow-intra-namespace NetworkPolicy to the namespace
  _, err := client.NetworkPolicies(namespace.Name).Create(intra)
  if err != nil {
    return err
  }

  return nil
}

func writeIntraPolicy(namespace *api.Namespace, routingLabelName string) *extensions.NetworkPolicy {
  return &extensions.NetworkPolicy{
    ObjectMeta: api.ObjectMeta{
      Name: IntraPolicyName, // name of the network policy
    },
    Spec: extensions.NetworkPolicySpec{
      PodSelector: unversionedApi.LabelSelector{}, // empty PodSelector selects all pods
      Ingress: []extensions.NetworkPolicyIngressRule{
        extensions.NetworkPolicyIngressRule{ // omit Ports allows on all ports
          From: []extensions.NetworkPolicyPeer{
            extensions.NetworkPolicyPeer{
              // allow pod traffic from this namespace
              NamespaceSelector: &unversionedApi.LabelSelector{
                MatchLabels: map[string]string{
                  routingLabelName: namespace.Name,
                },
              },
            },
          },
        },
      },
    },
  }
}

func writeBridgePolicy(config *utils.Config) *extensions.NetworkPolicy {
  return &extensions.NetworkPolicy{
    ObjectMeta: api.ObjectMeta{
      Name: BridgePolicyName, // name of the network policy
    },
    Spec: extensions.NetworkPolicySpec{
      PodSelector: unversionedApi.LabelSelector{}, // empty PodSelector selects all pods
      Ingress: []extensions.NetworkPolicyIngressRule{
        extensions.NetworkPolicyIngressRule{ // omit Ports allows on all ports
          From: []extensions.NetworkPolicyPeer{
            extensions.NetworkPolicyPeer{
              // allow pod traffic from the routing namespace
              NamespaceSelector: &unversionedApi.LabelSelector{
                MatchLabels: map[string]string{
                  config.RoutingLabelName: config.RoutingNamespace,
                },
              },
            },
          },
        },
      },
    },
  }
}