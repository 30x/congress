# congress

`congress` is a k8s Namespace watcher designed to utilize NetworkPolicy objects to create namespace isolation within a cluster.

**First, this is an Apigee specific implementation.**

All namespace creation events will be evaluated by `congress` and isolated if not specifically excluded. A namespace evaluated by `congress` needs to have an `annotation` that locks down all inter/intra-namespace pod traffic. It will be given a label that it's `NetworkPolicy` objects will use to identify the namespace's pod traffic. An example is below.
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
The namespace is also given two `NetworkPolicy` objects that behave as follows:
* `allow-intra-namespace`: this policy allows all pod traffic within the namespace (a.k.a "intra-namespace" traffic)
* `allow-apigee`: this policy allows pod traffic from the `apigee` namespace (a.k.a an "apigee bridge"). This is necessary for pods running on Apigee's cluster.

##Environment
There are two configurable environment variables:
* `CONGRESS_EXCLUDES`: a comma separated list of namespaces to exclude from network isolation (default: `"kube-system"`). Example: `"kube-system,default,myImportantNamespace,anotherImportantNamspace"`
* `CONGRESS_SELECTOR`: the label selector used by the watcher to filter `Namespace` events by (default: `""` i.e. all namespaces).
  Example: `"company=apigee,env=test"`
* `CONGRESS_IGNORE_SELECTOR`: a label selector used to suplement the excludes list in picking namespaces to ignore (default `""` i.e. ignore none)
