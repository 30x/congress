# congress

`congress` is a k8s Namespace watcher designed to utilize `NetworkPolicy` objects to create namespace isolation within a cluster.

**First, this is built to satisfy an Apigee use-case, but can easily be repurposed if you are in a similar situation.**

### When to use
`congress` is ideal for the cluster manager that has a dedicated pod router, such as the [k8s-router](https://github.com/30x/k8s-router), but also wants to control the internal pod-to-pod traffic.
Such a situation is commmon in multi-tenant Kubernetes clusters where it is undesirable to have `ClientA` pods capable of communication with `ClientB` pods.

### Design

`congress` primarily uses the [Kubernetes Namespace Watch API](https://godoc.org/k8s.io/kubernetes/pkg/client/unversioned#NamespaceInterface).
The main functionality can be broken into two parts: Namespace annotation/labelling and `NetworkPolicy` creation.

For the first piece of functionality, all namespace creation events will be evaluated by `congress` and isolated if not specifically excluded in either the configurable exclusion list or label. A namespace evaluated by `congress` needs to have an `annotation` that locks down all inter/intra-namespace pod traffic. It will be given a label that it's `NetworkPolicy` objects will use to identify the namespace's pod traffic. An example is below.
```yaml
kind: Namespace
apiVersion: v1
metadata:
  name: test
  labels:
   name: test # added by congress
  annotations:
    net.beta.kubernetes.io/network-policy: | # necessary for NetworkPolicies to work
      {
        "ingress": {
          "isolation": "DefaultDeny"
        }
      }
```
The idiomatic approach to this would be using an [AdmissionController](http://kubernetes.io/docs/admin/admission-controllers/). These are cumbersome to develop and maintain at the moment, so a dedicated Watcher pod is the next best approach.

For the second piece of functionality, upon creation the namespace is given two `NetworkPolicy` objects that behave as follows:
* `allow-intra-namespace`: this policy allows all pod traffic within the namespace (a.k.a "intra-namespace" traffic)
* `allow-routing`: this policy allows pod traffic from the namespace containing the pod router (a.k.a an "routing bridge").

##Environment
There are a few configurable (some required) environment variables:
* `CONGRESS_EXCLUDES`: a comma separated list of namespaces to exclude from network isolation (default: `"kube-system"`). Example: `"kube-system,default,myImportantNamespace,anotherImportantNamspace"`
* `CONGRESS_SELECTOR`: the label selector used by the watcher to filter `Namespace` events by (default: `""` i.e. all namespaces).
  Example: `"company=apigee,env=test"`
* `CONGRESS_IGNORE_SELECTOR`: a label selector used to suplement the excludes list in picking namespaces to ignore (default `""` i.e. ignore none)
* **required** `CONGRESS_ROUTING_NAMESPACE`: the namespace containing the pod router that forwards traffic
* **required** `CONGRESS_ROUTING_LABEL`: the key of the label used to identify pod traffic at the namespace level
* **required** `CONGRESS_ROUTING_POLICY_NAME`: the name of the `NetworkPolicy` that bridges the routing namespace to a tenant namespace
