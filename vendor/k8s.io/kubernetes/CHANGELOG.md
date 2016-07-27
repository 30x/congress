<!-- BEGIN MUNGE: GENERATED_TOC -->

- [v1.3.0-beta.2](#v130-beta2)
  - [Downloads](#downloads)
  - [Changes since v1.3.0-beta.1](#changes-since-v130-beta1)
    - [Experimental Features](#experimental-features)
    - [Other notable changes](#other-notable-changes)
- [v1.3.0-beta.1](#v130-beta1)
  - [Downloads](#downloads)
  - [Changes since v1.3.0-alpha.5](#changes-since-v130-alpha5)
    - [Action Required](#action-required)
    - [Other notable changes](#other-notable-changes)
- [v1.3.0-alpha.5](#v130-alpha5)
  - [Downloads](#downloads)
  - [Changes since v1.3.0-alpha.4](#changes-since-v130-alpha4)
    - [Action Required](#action-required)
    - [Other notable changes](#other-notable-changes)
- [v1.3.0-alpha.4](#v130-alpha4)
  - [Downloads](#downloads)
  - [Changes since v1.3.0-alpha.3](#changes-since-v130-alpha3)
    - [Action Required](#action-required)
    - [Other notable changes](#other-notable-changes)
- [v1.2.4](#v124)
  - [Downloads](#downloads)
  - [Changes since v1.2.3](#changes-since-v123)
    - [Other notable changes](#other-notable-changes)
- [v1.3.0-alpha.3](#v130-alpha3)
  - [Downloads](#downloads)
  - [Changes since v1.3.0-alpha.2](#changes-since-v130-alpha2)
    - [Action Required](#action-required)
    - [Other notable changes](#other-notable-changes)
- [v1.2.3](#v123)
  - [Downloads](#downloads)
  - [Changes since v1.2.2](#changes-since-v122)
    - [Action Required](#action-required)
    - [Other notable changes](#other-notable-changes)
- [v1.3.0-alpha.2](#v130-alpha2)
  - [Downloads](#downloads)
  - [Changes since v1.3.0-alpha.1](#changes-since-v130-alpha1)
    - [Other notable changes](#other-notable-changes)
- [v1.2.2](#v122)
  - [Downloads](#downloads)
  - [Changes since v1.2.1](#changes-since-v121)
    - [Other notable changes](#other-notable-changes)
- [v1.2.1](#v121)
  - [Downloads](#downloads)
  - [Changes since v1.2.0](#changes-since-v120)
    - [Other notable changes](#other-notable-changes)
- [v1.3.0-alpha.1](#v130-alpha1)
  - [Downloads](#downloads)
  - [Changes since v1.2.0](#changes-since-v120)
    - [Action Required](#action-required)
    - [Other notable changes](#other-notable-changes)
- [v1.2.0](#v120)
  - [Downloads](#downloads)
  - [Changes since v1.1.1](#changes-since-v111)
    - [Major Themes](#major-themes)
    - [Other notable improvements](#other-notable-improvements)
    - [Experimental Features](#experimental-features)
    - [Action required](#action-required)
    - [Known Issues](#known-issues)
      - [Docker Known Issues](#docker-known-issues)
        - [1.9.1](#191)
    - [Provider-specific Notes](#provider-specific-notes)
      - [Various](#various)
      - [AWS](#aws)
      - [GCE](#gce)

<!-- END MUNGE: GENERATED_TOC -->

<!-- NEW RELEASE NOTES ENTRY -->


# v1.3.0-beta.2

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/release-1.3/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.3.0-beta.2/kubernetes.tar.gz) | `9c95762970b943d6c6547f0841c1e5471148b0e3` | `dc9e8560f24459b2313317b15910bee7`

## Changes since v1.3.0-beta.1

### Experimental Features

* Proposal for implementing init containers ([#23666](https://github.com/kubernetes/kubernetes/pull/23666), [@smarterclayton](https://github.com/smarterclayton))

### Other notable changes

* GCE provider: Limit Filter calls to regexps rather than insane blobs ([#27741](https://github.com/kubernetes/kubernetes/pull/27741), [@zmerlynn](https://github.com/zmerlynn))
* swap FIRSTSEEN/LASTSEEN columns in `kubectl get event -w` ([#27549](https://github.com/kubernetes/kubernetes/pull/27549), [@therc](https://github.com/therc))
* GCI: fix kubectl permission issue [#27643](https://github.com/kubernetes/kubernetes/pull/27643) ([#27740](https://github.com/kubernetes/kubernetes/pull/27740), [@andyzheng0831](https://github.com/andyzheng0831))
* Add federation api and cm servers to hyperkube ([#27586](https://github.com/kubernetes/kubernetes/pull/27586), [@colhom](https://github.com/colhom))
* federation: Creating kubeconfig files to be used for creating secrets for clusters on aws and gke ([#27332](https://github.com/kubernetes/kubernetes/pull/27332), [@nikhiljindal](https://github.com/nikhiljindal))
* AWS: Enable ICMP Type 3 Code 4 for ELBs ([#27677](https://github.com/kubernetes/kubernetes/pull/27677), [@justinsb](https://github.com/justinsb))
* Bumped Heapster to v1.1.0 ([#27542](https://github.com/kubernetes/kubernetes/pull/27542), [@piosz](https://github.com/piosz))
* Deleting federation-push.sh ([#27400](https://github.com/kubernetes/kubernetes/pull/27400), [@nikhiljindal](https://github.com/nikhiljindal))
* Validate-cluster finishes shortly after at most ALLOWED_NOTREADY_NODE… ([#26778](https://github.com/kubernetes/kubernetes/pull/26778), [@gmarek](https://github.com/gmarek))
* AWS kube-down: Issue warning if VPC not found ([#27518](https://github.com/kubernetes/kubernetes/pull/27518), [@justinsb](https://github.com/justinsb))
* gce/kube-down: Parallelize IGM deletion, batch more ([#27302](https://github.com/kubernetes/kubernetes/pull/27302), [@zmerlynn](https://github.com/zmerlynn))
* Enable dynamic allocation of heapster/eventer cpu request/limit ([#27185](https://github.com/kubernetes/kubernetes/pull/27185), [@gmarek](https://github.com/gmarek))
* 'kubectl describe pv' now shows events ([#27431](https://github.com/kubernetes/kubernetes/pull/27431), [@jsafrane](https://github.com/jsafrane))
* AWS kube-up: set net.ipv4.neigh.default.gc_thresh1=0 to avoid ARP over-caching ([#27682](https://github.com/kubernetes/kubernetes/pull/27682), [@justinsb](https://github.com/justinsb))
* AWS volumes: Use /dev/xvdXX names with EC2 ([#27628](https://github.com/kubernetes/kubernetes/pull/27628), [@justinsb](https://github.com/justinsb))
* Prep for continuous Docker validation test ([#26813](https://github.com/kubernetes/kubernetes/pull/26813), [@wonderfly](https://github.com/wonderfly))
* Bump cAdvisor to v0.23.4 ([#27591](https://github.com/kubernetes/kubernetes/pull/27591), [@dchen1107](https://github.com/dchen1107))
* Change default value of deleting-pods-burst to 1 ([#27606](https://github.com/kubernetes/kubernetes/pull/27606), [@gmarek](https://github.com/gmarek))
* MESOS: fix race condition in contrib/mesos/pkg/queue/delay ([#24916](https://github.com/kubernetes/kubernetes/pull/24916), [@jdef](https://github.com/jdef))
* including federation binaries in the list of images we push during release ([#27396](https://github.com/kubernetes/kubernetes/pull/27396), [@nikhiljindal](https://github.com/nikhiljindal))
* fix updatePod() of RS and RC controllers ([#27415](https://github.com/kubernetes/kubernetes/pull/27415), [@caesarxuchao](https://github.com/caesarxuchao))
* Change default value of deleting-pods-burst to 1 ([#27422](https://github.com/kubernetes/kubernetes/pull/27422), [@gmarek](https://github.com/gmarek))
* Kubelet Volume Attach/Detach/Mount/Unmount Redesign ([#26801](https://github.com/kubernetes/kubernetes/pull/26801), [@saad-ali](https://github.com/saad-ali))



# v1.3.0-beta.1

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/release-1.3/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.3.0-beta.1/kubernetes.tar.gz) | `2b54995ee8f52d78dc31c3d7291e8dfa5c809fe7` | `f1022a84c3441cae4ebe1d295470be8f`

## Changes since v1.3.0-alpha.5

### Action Required

* Fixing logic to generate ExternalHost in genericapiserver ([#26796](https://github.com/kubernetes/kubernetes/pull/26796), [@nikhiljindal](https://github.com/nikhiljindal))
* federation: Updating federation-controller-manager to use secret to get federation-apiserver's kubeconfig ([#26819](https://github.com/kubernetes/kubernetes/pull/26819), [@nikhiljindal](https://github.com/nikhiljindal))

### Other notable changes

* federation: fix dns provider initialization issues ([#27252](https://github.com/kubernetes/kubernetes/pull/27252), [@mfanjie](https://github.com/mfanjie))
* Updating federation up scripts to work in non e2e setup ([#27260](https://github.com/kubernetes/kubernetes/pull/27260), [@nikhiljindal](https://github.com/nikhiljindal))
* version bump for gci to milestone 53 ([#27210](https://github.com/kubernetes/kubernetes/pull/27210), [@adityakali](https://github.com/adityakali))
* kubectl apply retry stale resource version ([#26557](https://github.com/kubernetes/kubernetes/pull/26557), [@AdoHe](https://github.com/AdoHe))
* Revert "Wait for arc.getArchive() to complete before running tests" ([#27130](https://github.com/kubernetes/kubernetes/pull/27130), [@pwittrock](https://github.com/pwittrock))
* ResourceQuota BestEffort scope aligned with Pod level QoS ([#26969](https://github.com/kubernetes/kubernetes/pull/26969), [@derekwaynecarr](https://github.com/derekwaynecarr))
* AWS: cache instances during service reload to avoid rate limiting on restart ([#26900](https://github.com/kubernetes/kubernetes/pull/26900), [@therc](https://github.com/therc))
* GCE provider: Log full contents of long operations ([#26962](https://github.com/kubernetes/kubernetes/pull/26962), [@zmerlynn](https://github.com/zmerlynn))
* Fix system container detection ([#26586](https://github.com/kubernetes/kubernetes/pull/26586), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Added hpa/v1 generator to kubectl autoscale ([#26775](https://github.com/kubernetes/kubernetes/pull/26775), [@piosz](https://github.com/piosz))
* federation: Adding dnsprovider flags to federation-controller-manager ([#27158](https://github.com/kubernetes/kubernetes/pull/27158), [@nikhiljindal](https://github.com/nikhiljindal))
* federation service controller: fixing a bug so that existing services are created in newly registered clusters ([#27028](https://github.com/kubernetes/kubernetes/pull/27028), [@mfanjie](https://github.com/mfanjie))
* Rename ENABLE_NODE_AUTOSCALER to ENABLE_CLUSTER_AUTOSCALER - part 2 ([#27117](https://github.com/kubernetes/kubernetes/pull/27117), [@mwielgus](https://github.com/mwielgus))
* support for mounting local-ssds on GCI ([#27143](https://github.com/kubernetes/kubernetes/pull/27143), [@adityakali](https://github.com/adityakali))
* AWS: support mixed plaintext/encrypted ports in ELBs via service.beta.kubernetes.io/aws-load-balancer-ssl-ports annotation ([#26976](https://github.com/kubernetes/kubernetes/pull/26976), [@therc](https://github.com/therc))
* Updating e2e docs with instructions on running federation tests ([#27072](https://github.com/kubernetes/kubernetes/pull/27072), [@colhom](https://github.com/colhom))
* LBaaS v2 Support for Openstack Cloud Provider Plugin ([#25987](https://github.com/kubernetes/kubernetes/pull/25987), [@dagnello](https://github.com/dagnello))
* GCI: add support for network plugin ([#27027](https://github.com/kubernetes/kubernetes/pull/27027), [@andyzheng0831](https://github.com/andyzheng0831))
* Bump cAdvisor to v0.23.3 ([#27065](https://github.com/kubernetes/kubernetes/pull/27065), [@timstclair](https://github.com/timstclair))
* Stop 'kubectl drain' deleting pods with local storage. ([#26667](https://github.com/kubernetes/kubernetes/pull/26667), [@mml](https://github.com/mml))
* Networking e2es: Wait for all nodes to be schedulable in kubeproxy and networking tests ([#27008](https://github.com/kubernetes/kubernetes/pull/27008), [@zmerlynn](https://github.com/zmerlynn))
* change clientset of service controller to versioned ([#26694](https://github.com/kubernetes/kubernetes/pull/26694), [@mfanjie](https://github.com/mfanjie))
* GCE: Enable using gcr.io as a Docker registry mirror. ([#25841](https://github.com/kubernetes/kubernetes/pull/25841), [@ojarjur](https://github.com/ojarjur))
* correction on rbd volume object and defaults ([#25490](https://github.com/kubernetes/kubernetes/pull/25490), [@rootfs](https://github.com/rootfs))
* Bump GCE debian image to container-v1-3-v20160604 ([#26851](https://github.com/kubernetes/kubernetes/pull/26851), [@zmerlynn](https://github.com/zmerlynn))
* Option to enable http2 on client connections. ([#25280](https://github.com/kubernetes/kubernetes/pull/25280), [@timothysc](https://github.com/timothysc))
* kubectl get ingress output remove rules ([#26684](https://github.com/kubernetes/kubernetes/pull/26684), [@AdoHe](https://github.com/AdoHe))
* AWS kube-up: Remove SecurityContextDeny admission controller (to mirror GCE) ([#25381](https://github.com/kubernetes/kubernetes/pull/25381), [@zquestz](https://github.com/zquestz))
* Fix third party ([#25894](https://github.com/kubernetes/kubernetes/pull/25894), [@brendandburns](https://github.com/brendandburns))
* AWS Route53 dnsprovider ([#26049](https://github.com/kubernetes/kubernetes/pull/26049), [@quinton-hoole](https://github.com/quinton-hoole))
* GCI/Trusty: support the Docker registry mirror ([#26745](https://github.com/kubernetes/kubernetes/pull/26745), [@andyzheng0831](https://github.com/andyzheng0831))
* Attach/Detach Controller Kubelet Changes ([#26351](https://github.com/kubernetes/kubernetes/pull/26351), [@saad-ali](https://github.com/saad-ali))
* Added DNS Reverse Record logic for service IPs ([#26226](https://github.com/kubernetes/kubernetes/pull/26226), [@ArtfulCoder](https://github.com/ArtfulCoder))
* read gluster log to surface glusterfs plugin errors properly in describe events ([#24808](https://github.com/kubernetes/kubernetes/pull/24808), [@screeley44](https://github.com/screeley44))



# v1.3.0-alpha.5

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/master/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.3.0-alpha.5/kubernetes.tar.gz) | `724bf5a4437ca9dc75d9297382f47a179e8dc5a6` | `2a8b4a5297df3007fce69f1e344fd87e`

## Changes since v1.3.0-alpha.4

### Action Required

* Add direct serializer ([#26251](https://github.com/kubernetes/kubernetes/pull/26251), [@caesarxuchao](https://github.com/caesarxuchao))
* Add a NodeCondition "NetworkUnavailable" to prevent scheduling onto a node until the routes have been created  ([#26415](https://github.com/kubernetes/kubernetes/pull/26415), [@wojtek-t](https://github.com/wojtek-t))
* Add garbage collector into kube-controller-manager ([#26341](https://github.com/kubernetes/kubernetes/pull/26341), [@caesarxuchao](https://github.com/caesarxuchao))
* Add orphaning finalizer logic to GC ([#25599](https://github.com/kubernetes/kubernetes/pull/25599), [@caesarxuchao](https://github.com/caesarxuchao))
* GCI-backed masters mount srv/kubernetes and srv/sshproxy in the right place ([#26238](https://github.com/kubernetes/kubernetes/pull/26238), [@ihmccreery](https://github.com/ihmccreery))
* Updaing QoS policy to be at the pod level ([#14943](https://github.com/kubernetes/kubernetes/pull/14943), [@vishh](https://github.com/vishh))
* add CIDR allocator for NodeController ([#19242](https://github.com/kubernetes/kubernetes/pull/19242), [@mqliang](https://github.com/mqliang))
* Adding garbage collector controller ([#24509](https://github.com/kubernetes/kubernetes/pull/24509), [@caesarxuchao](https://github.com/caesarxuchao))

### Other notable changes

* Fix a bug with pluralization of third party resources ([#25374](https://github.com/kubernetes/kubernetes/pull/25374), [@brendandburns](https://github.com/brendandburns))
* Run l7 controller on master  ([#26048](https://github.com/kubernetes/kubernetes/pull/26048), [@bprashanth](https://github.com/bprashanth))
* AWS: ELB proxy protocol support via annotation service.beta.kubernetes.io/aws-load-balancer-proxy-protocol ([#24569](https://github.com/kubernetes/kubernetes/pull/24569), [@williamsandrew](https://github.com/williamsandrew))
* kubectl run --restart=Never creates pods ([#25253](https://github.com/kubernetes/kubernetes/pull/25253), [@soltysh](https://github.com/soltysh))
* Add LabelSelector to PersistentVolumeClaimSpec ([#25917](https://github.com/kubernetes/kubernetes/pull/25917), [@pmorie](https://github.com/pmorie))
* Removed metrics api group ([#26073](https://github.com/kubernetes/kubernetes/pull/26073), [@piosz](https://github.com/piosz))
* Fixed check in kubectl autoscale. ([#26162](https://github.com/kubernetes/kubernetes/pull/26162), [@jszczepkowski](https://github.com/jszczepkowski))
* Add support for 3rd party objects to kubectl label ([#24882](https://github.com/kubernetes/kubernetes/pull/24882), [@brendandburns](https://github.com/brendandburns))
* Move shell completion generation into 'kubectl completion' command ([#23801](https://github.com/kubernetes/kubernetes/pull/23801), [@sttts](https://github.com/sttts))
* Setting TLS1.2 minimum because TLS1.0 and TLS1.1 are vulnerable ([#26169](https://github.com/kubernetes/kubernetes/pull/26169), [@victorgp](https://github.com/victorgp))
* Kubelet: Periodically reporting image pulling progress in log ([#26145](https://github.com/kubernetes/kubernetes/pull/26145), [@Random-Liu](https://github.com/Random-Liu))
* Federation service controller ([#26034](https://github.com/kubernetes/kubernetes/pull/26034), [@mfanjie](https://github.com/mfanjie))
* Stabilize map order in kubectl describe ([#26046](https://github.com/kubernetes/kubernetes/pull/26046), [@timoreimann](https://github.com/timoreimann))
* Google Cloud DNS dnsprovider - replacement for [#25389](https://github.com/kubernetes/kubernetes/pull/25389) ([#26020](https://github.com/kubernetes/kubernetes/pull/26020), [@quinton-hoole](https://github.com/quinton-hoole))
* Fix system container detection in kubelet on systemd ([#25982](https://github.com/kubernetes/kubernetes/pull/25982), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Added pods-per-core to kubelet. [#25762](https://github.com/kubernetes/kubernetes/pull/25762) ([#25813](https://github.com/kubernetes/kubernetes/pull/25813), [@rrati](https://github.com/rrati))
* promote sourceRange into service spec ([#25826](https://github.com/kubernetes/kubernetes/pull/25826), [@freehan](https://github.com/freehan))
* kube-controller-manager: Add configure-cloud-routes option ([#25614](https://github.com/kubernetes/kubernetes/pull/25614), [@justinsb](https://github.com/justinsb))
* kubelet: reading cloudinfo from cadvisor ([#21373](https://github.com/kubernetes/kubernetes/pull/21373), [@enoodle](https://github.com/enoodle))
* Disable cAdvisor event storage by default ([#24771](https://github.com/kubernetes/kubernetes/pull/24771), [@timstclair](https://github.com/timstclair))
* Remove docker-multinode ([#26031](https://github.com/kubernetes/kubernetes/pull/26031), [@luxas](https://github.com/luxas))
* nodecontroller: Fix log message on successful update ([#26207](https://github.com/kubernetes/kubernetes/pull/26207), [@zmerlynn](https://github.com/zmerlynn))
* remove deprecated generated typed clients ([#26336](https://github.com/kubernetes/kubernetes/pull/26336), [@caesarxuchao](https://github.com/caesarxuchao))
* Kubenet host-port support through iptables ([#25604](https://github.com/kubernetes/kubernetes/pull/25604), [@freehan](https://github.com/freehan))
* Add metrics support for a GCE PD, EC2 EBS & Azure File volumes ([#25852](https://github.com/kubernetes/kubernetes/pull/25852), [@vishh](https://github.com/vishh))
* Bump cAdvisor (and dependencies) godeps version ([#25914](https://github.com/kubernetes/kubernetes/pull/25914), [@timstclair](https://github.com/timstclair))
* Add RBAC authorization API group and authorizer ([#25634](https://github.com/kubernetes/kubernetes/pull/25634), [@ericchiang](https://github.com/ericchiang))
* Add Seccomp to Annotations ([#25324](https://github.com/kubernetes/kubernetes/pull/25324), [@jfrazelle](https://github.com/jfrazelle))
* AWS: Fix long-standing bug in stringSetToPointers ([#26331](https://github.com/kubernetes/kubernetes/pull/26331), [@therc](https://github.com/therc))
* Add dnsmasq as a DNS cache in kube-dns pod ([#26114](https://github.com/kubernetes/kubernetes/pull/26114), [@ArtfulCoder](https://github.com/ArtfulCoder))
* routecontroller: Add wait.NonSlidingUntil, use it ([#26301](https://github.com/kubernetes/kubernetes/pull/26301), [@zmerlynn](https://github.com/zmerlynn))
* Attempt 2: Bump GCE containerVM to container-v1-3-v20160517 (Docker 1.11.1) again. ([#26001](https://github.com/kubernetes/kubernetes/pull/26001), [@dchen1107](https://github.com/dchen1107))
* Downward API implementation for resources limits and requests ([#24179](https://github.com/kubernetes/kubernetes/pull/24179), [@aveshagarwal](https://github.com/aveshagarwal))
* Replace containervm with GCI as default master image for GCE clusters ([#26197](https://github.com/kubernetes/kubernetes/pull/26197), [@wonderfly](https://github.com/wonderfly))
* Add a 'kubectl clusterinfo dump' option ([#20672](https://github.com/kubernetes/kubernetes/pull/20672), [@brendandburns](https://github.com/brendandburns))
* Fixing heapster memory requirements. ([#26109](https://github.com/kubernetes/kubernetes/pull/26109), [@Q-Lee](https://github.com/Q-Lee))
* Handle federated service name lookups in kube-dns. ([#25727](https://github.com/kubernetes/kubernetes/pull/25727), [@madhusudancs](https://github.com/madhusudancs))
* Support sort-by timestamp in kubectl get ([#25600](https://github.com/kubernetes/kubernetes/pull/25600), [@janetkuo](https://github.com/janetkuo))
* vSphere Volume Plugin Implementation ([#24947](https://github.com/kubernetes/kubernetes/pull/24947), [@abithap](https://github.com/abithap))
* ResourceQuota controller uses rate limiter to prevent hot-loops in error situations ([#25748](https://github.com/kubernetes/kubernetes/pull/25748), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Fix hyperkube flag parsing ([#25512](https://github.com/kubernetes/kubernetes/pull/25512), [@colhom](https://github.com/colhom))
* Add a kubectl create secret tls command ([#24719](https://github.com/kubernetes/kubernetes/pull/24719), [@bprashanth](https://github.com/bprashanth))
* Add node problem detector as an addon pod. ([#25986](https://github.com/kubernetes/kubernetes/pull/25986), [@Random-Liu](https://github.com/Random-Liu))
* Handle cAdvisor partial failures ([#25933](https://github.com/kubernetes/kubernetes/pull/25933), [@timstclair](https://github.com/timstclair))
* Use SkyDNS as a library for a more integrated kube DNS ([#23930](https://github.com/kubernetes/kubernetes/pull/23930), [@ArtfulCoder](https://github.com/ArtfulCoder))
* Introduce node memory pressure condition to scheduler ([#25531](https://github.com/kubernetes/kubernetes/pull/25531), [@ingvagabund](https://github.com/ingvagabund))
* Fix detection of docker cgroup on RHEL ([#25907](https://github.com/kubernetes/kubernetes/pull/25907), [@ncdc](https://github.com/ncdc))
* Add support for limiting grace period during soft eviction ([#25772](https://github.com/kubernetes/kubernetes/pull/25772), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Use protobufs by default to communicate with apiserver (still store JSONs in etcd) ([#25738](https://github.com/kubernetes/kubernetes/pull/25738), [@wojtek-t](https://github.com/wojtek-t))
* Add NetworkPolicy API Resource ([#25638](https://github.com/kubernetes/kubernetes/pull/25638), [@caseydavenport](https://github.com/caseydavenport))
* Only expose top N images in `NodeStatus` ([#25328](https://github.com/kubernetes/kubernetes/pull/25328), [@resouer](https://github.com/resouer))
* Extend secrets volumes with path control ([#25285](https://github.com/kubernetes/kubernetes/pull/25285), [@ingvagabund](https://github.com/ingvagabund))
* Implement OIDC client AuthProvider ([#25270](https://github.com/kubernetes/kubernetes/pull/25270), [@bobbyrullo](https://github.com/bobbyrullo))
* Make addon-manager cross-platform and use it with hyperkube ([#25631](https://github.com/kubernetes/kubernetes/pull/25631), [@luxas](https://github.com/luxas))
* kubelet: Optionally, have kubelet exit if lock file contention is observed, using --exit-on-lock-contention flag ([#25596](https://github.com/kubernetes/kubernetes/pull/25596), [@derekparker](https://github.com/derekparker))
* Bump up glbc version to 0.6.2 ([#25446](https://github.com/kubernetes/kubernetes/pull/25446), [@bprashanth](https://github.com/bprashanth))
* Add 'kubectl set image' ([#25509](https://github.com/kubernetes/kubernetes/pull/25509), [@janetkuo](https://github.com/janetkuo))
* NodeController doesn't evict Pods if no Nodes are Ready ([#25571](https://github.com/kubernetes/kubernetes/pull/25571), [@gmarek](https://github.com/gmarek))
* Added enforcing of setting nodes numbers for cluster autoscaler. ([#25734](https://github.com/kubernetes/kubernetes/pull/25734), [@jszczepkowski](https://github.com/jszczepkowski))
* systemd node spec proposal ([#17688](https://github.com/kubernetes/kubernetes/pull/17688), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Bump GCE ContainerVM to container-v1-3-v20160517 (Docker 1.11.1) ([#25843](https://github.com/kubernetes/kubernetes/pull/25843), [@zmerlynn](https://github.com/zmerlynn))
* AWS: Move enforcement of attached AWS device limit from kubelet to scheduler ([#23254](https://github.com/kubernetes/kubernetes/pull/23254), [@jsafrane](https://github.com/jsafrane))
* Refactor persistent volume controller ([#24331](https://github.com/kubernetes/kubernetes/pull/24331), [@jsafrane](https://github.com/jsafrane))
* Add support for running GCI on the GCE cloud provider ([#25425](https://github.com/kubernetes/kubernetes/pull/25425), [@andyzheng0831](https://github.com/andyzheng0831))
* Implement taints and tolerations ([#24134](https://github.com/kubernetes/kubernetes/pull/24134), [@kevin-wangzefeng](https://github.com/kevin-wangzefeng))
* GCI: Ensure that the right version of kubelet is used ([#25504](https://github.com/kubernetes/kubernetes/pull/25504), [@andyzheng0831](https://github.com/andyzheng0831))
* Add init containers to pods ([#23567](https://github.com/kubernetes/kubernetes/pull/23567), [@smarterclayton](https://github.com/smarterclayton))
* GCI: Fix the condition for using the default image ([#25763](https://github.com/kubernetes/kubernetes/pull/25763), [@andyzheng0831](https://github.com/andyzheng0831))



# v1.3.0-alpha.4

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/master/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.3.0-alpha.4/kubernetes.tar.gz) | `758e97e7e50153840379ecd9f8fda1869543539f` | `4e18ae6a428c99fcc30e2137d7c41854`

## Changes since v1.3.0-alpha.3

### Action Required

* validate third party resources ([#25007](https://github.com/kubernetes/kubernetes/pull/25007), [@liggitt](https://github.com/liggitt))
* Automatically create the kube-system namespace ([#25196](https://github.com/kubernetes/kubernetes/pull/25196), [@luxas](https://github.com/luxas))
* Make ThirdPartyResource a root scoped object ([#25006](https://github.com/kubernetes/kubernetes/pull/25006), [@liggitt](https://github.com/liggitt))
* mark container-port flag as deprecated ([#25072](https://github.com/kubernetes/kubernetes/pull/25072), [@AdoHe](https://github.com/AdoHe))
* Provide flags to use etcd3 backed storage ([#24455](https://github.com/kubernetes/kubernetes/pull/24455), [@hongchaodeng](https://github.com/hongchaodeng))

### Other notable changes

* Fix hyperkube's layer caching, and remove --make-symlinks at build time ([#25693](https://github.com/kubernetes/kubernetes/pull/25693), [@luxas](https://github.com/luxas))
* AWS: More support for ap-northeast-2 region ([#24464](https://github.com/kubernetes/kubernetes/pull/24464), [@matthewrudy](https://github.com/matthewrudy))
* Make bigger master root disks in GCE for large clusters ([#25670](https://github.com/kubernetes/kubernetes/pull/25670), [@gmarek](https://github.com/gmarek))
* AWS kube-down: don't fail if ELB not in VPC - [#23784](https://github.com/kubernetes/kubernetes/pull/23784) ([#23785](https://github.com/kubernetes/kubernetes/pull/23785), [@ajohnstone](https://github.com/ajohnstone))
* Build hyperkube in hack/local-up-cluster instead of separate binaries ([#25627](https://github.com/kubernetes/kubernetes/pull/25627), [@luxas](https://github.com/luxas))
* enable recursive processing in kubectl rollout ([#25110](https://github.com/kubernetes/kubernetes/pull/25110), [@metral](https://github.com/metral))
* Support struct,array,slice types when sorting kubectl output ([#25022](https://github.com/kubernetes/kubernetes/pull/25022), [@zhouhaibing089](https://github.com/zhouhaibing089))
* federated api servers: Adding a discovery summarizer server ([#20358](https://github.com/kubernetes/kubernetes/pull/20358), [@nikhiljindal](https://github.com/nikhiljindal))
* AWS: Allow cross-region image pulling with ECR ([#24369](https://github.com/kubernetes/kubernetes/pull/24369), [@therc](https://github.com/therc))
* Automatically add node labels beta.kubernetes.io/{os,arch} ([#23684](https://github.com/kubernetes/kubernetes/pull/23684), [@luxas](https://github.com/luxas))
* kubectl suggest for get (list, ps), and delete(rm) ([#25181](https://github.com/kubernetes/kubernetes/pull/25181), [@janetkuo](https://github.com/janetkuo))
* Add IPv6 address support for pods - does NOT include services ([#23090](https://github.com/kubernetes/kubernetes/pull/23090), [@tgraf](https://github.com/tgraf))
* Use local disk for ConfigMap volume instead of tmpfs ([#25306](https://github.com/kubernetes/kubernetes/pull/25306), [@pmorie](https://github.com/pmorie))
* NVIDIA GPU support ([#24836](https://github.com/kubernetes/kubernetes/pull/24836), [@therc](https://github.com/therc))
* AWS: SSL support for ELB listeners through annotations ([#23495](https://github.com/kubernetes/kubernetes/pull/23495), [@therc](https://github.com/therc))
* Add `kubectl rollout status` ([#19946](https://github.com/kubernetes/kubernetes/pull/19946), [@janetkuo](https://github.com/janetkuo))
* Webhook Token Authenticator ([#24902](https://github.com/kubernetes/kubernetes/pull/24902), [@cjcullen](https://github.com/cjcullen))
* PSP admission ([#24600](https://github.com/kubernetes/kubernetes/pull/24600), [@pweil-](https://github.com/pweil-))
* Scheduledjob api ([#24970](https://github.com/kubernetes/kubernetes/pull/24970), [@soltysh](https://github.com/soltysh))
* Kubectl support for validating nested objects with different ApiGroups ([#25172](https://github.com/kubernetes/kubernetes/pull/25172), [@pwittrock](https://github.com/pwittrock))
* Change default clusterCIDRs from /16 to /14 in GCE configs allowing 1000 Node clusters by default. ([#25350](https://github.com/kubernetes/kubernetes/pull/25350), [@gmarek](https://github.com/gmarek))
* Add 'kubectl set' ([#25444](https://github.com/kubernetes/kubernetes/pull/25444), [@janetkuo](https://github.com/janetkuo))
* vSphere Cloud Provider Implementation  ([#24703](https://github.com/kubernetes/kubernetes/pull/24703), [@dagnello](https://github.com/dagnello))
* Added JobTemplate, a preliminary step for ScheduledJob and Workflow ([#21675](https://github.com/kubernetes/kubernetes/pull/21675), [@soltysh](https://github.com/soltysh))
* Openstack provider ([#21737](https://github.com/kubernetes/kubernetes/pull/21737), [@zreigz](https://github.com/zreigz))
* AWS kube-up: Allow VPC CIDR to be specified (experimental) ([#23362](https://github.com/kubernetes/kubernetes/pull/23362), [@miguelfrde](https://github.com/miguelfrde))
* Return "410 Gone" errors via watch stream when using watch cache ([#25369](https://github.com/kubernetes/kubernetes/pull/25369), [@liggitt](https://github.com/liggitt))
* GKE provider: Add cluster-ipv4-cidr and arbitrary flags ([#25437](https://github.com/kubernetes/kubernetes/pull/25437), [@zmerlynn](https://github.com/zmerlynn))
* AWS kube-up: Increase timeout waiting for docker start ([#25405](https://github.com/kubernetes/kubernetes/pull/25405), [@justinsb](https://github.com/justinsb))
* Sort resources in quota errors to avoid duplicate events ([#25161](https://github.com/kubernetes/kubernetes/pull/25161), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Display line number on JSON errors ([#25038](https://github.com/kubernetes/kubernetes/pull/25038), [@mfojtik](https://github.com/mfojtik))
* GCE: Allow node count to exceed GCE TargetPool maximums ([#25178](https://github.com/kubernetes/kubernetes/pull/25178), [@zmerlynn](https://github.com/zmerlynn))
* Clarify supported version skew between masters, nodes, and clients ([#25087](https://github.com/kubernetes/kubernetes/pull/25087), [@ihmccreery](https://github.com/ihmccreery))
* Move godeps to vendor/ ([#24242](https://github.com/kubernetes/kubernetes/pull/24242), [@thockin](https://github.com/thockin))
* Introduce events flag for describers ([#24554](https://github.com/kubernetes/kubernetes/pull/24554), [@ingvagabund](https://github.com/ingvagabund))
* run kube-addon-manager in a static pod ([#23600](https://github.com/kubernetes/kubernetes/pull/23600), [@mikedanese](https://github.com/mikedanese))
* Reimplement 'pause' in C - smaller footprint all around ([#23009](https://github.com/kubernetes/kubernetes/pull/23009), [@uluyol](https://github.com/uluyol))
* Add subPath to mount a child dir or file of a volumeMount ([#22575](https://github.com/kubernetes/kubernetes/pull/22575), [@MikaelCluseau](https://github.com/MikaelCluseau))
* Handle image digests in node status and image GC ([#25088](https://github.com/kubernetes/kubernetes/pull/25088), [@ncdc](https://github.com/ncdc))
* PLEG: reinspect pods that failed prior inspections ([#25077](https://github.com/kubernetes/kubernetes/pull/25077), [@ncdc](https://github.com/ncdc))
* Upgrade installed packages when building hyperkube to improve the security profile ([#25114](https://github.com/kubernetes/kubernetes/pull/25114), [@aaronlevy](https://github.com/aaronlevy))
* GCI/Trusty: Support ABAC authorization ([#24950](https://github.com/kubernetes/kubernetes/pull/24950), [@andyzheng0831](https://github.com/andyzheng0831))
* fix cinder volume dir umount issue [#24717](https://github.com/kubernetes/kubernetes/pull/24717) ([#24718](https://github.com/kubernetes/kubernetes/pull/24718), [@chengyli](https://github.com/chengyli))
* Inter pod topological affinity and anti-affinity implementation ([#22985](https://github.com/kubernetes/kubernetes/pull/22985), [@kevin-wangzefeng](https://github.com/kevin-wangzefeng))
* start etcd compactor in background ([#25010](https://github.com/kubernetes/kubernetes/pull/25010), [@hongchaodeng](https://github.com/hongchaodeng))
* GCI: Add two GCI specific metadata pairs ([#25105](https://github.com/kubernetes/kubernetes/pull/25105), [@andyzheng0831](https://github.com/andyzheng0831))
* Ensure status is not changed during an update of PV, PVC, HPA objects ([#24924](https://github.com/kubernetes/kubernetes/pull/24924), [@mqliang](https://github.com/mqliang))
* GCE: Prefer preconfigured node tags for firewalls, if available ([#25148](https://github.com/kubernetes/kubernetes/pull/25148), [@a-robinson](https://github.com/a-robinson))
* kubectl rolling-update support for same image ([#24645](https://github.com/kubernetes/kubernetes/pull/24645), [@jlowdermilk](https://github.com/jlowdermilk))
* Update salt config to allow Debian Jessie on GCE. ([#25123](https://github.com/kubernetes/kubernetes/pull/25123), [@jlewi](https://github.com/jlewi))
* Mark kube-push.sh as broken ([#25095](https://github.com/kubernetes/kubernetes/pull/25095), [@ihmccreery](https://github.com/ihmccreery))
* AWS: Add support for ap-northeast-2 region (Seoul) ([#24457](https://github.com/kubernetes/kubernetes/pull/24457), [@leokhoa](https://github.com/leokhoa))
* GCI: Update the command to get the image ([#24987](https://github.com/kubernetes/kubernetes/pull/24987), [@andyzheng0831](https://github.com/andyzheng0831))
* Port-forward: use out and error streams instead of glog ([#17030](https://github.com/kubernetes/kubernetes/pull/17030), [@csrwng](https://github.com/csrwng))
* Promote Pod Hostname & Subdomain to fields (were annotations) ([#24362](https://github.com/kubernetes/kubernetes/pull/24362), [@ArtfulCoder](https://github.com/ArtfulCoder))
* Validate deletion timestamp doesn't change on update ([#24839](https://github.com/kubernetes/kubernetes/pull/24839), [@liggitt](https://github.com/liggitt))
* Add flag -t as shorthand for --tty ([#24365](https://github.com/kubernetes/kubernetes/pull/24365), [@janetkuo](https://github.com/janetkuo))
* Add support for running clusters on GCI ([#24893](https://github.com/kubernetes/kubernetes/pull/24893), [@andyzheng0831](https://github.com/andyzheng0831))
* Switch to ABAC authorization from AllowAll ([#24210](https://github.com/kubernetes/kubernetes/pull/24210), [@cjcullen](https://github.com/cjcullen))
* Fix DeletingLoadBalancer event generation. ([#24833](https://github.com/kubernetes/kubernetes/pull/24833), [@a-robinson](https://github.com/a-robinson))



# v1.2.4

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/release-1.2/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.2.4/kubernetes.tar.gz) | `f3aea83f8f0e16b2b41998a2edc09eb42fd8d945` | `ab0aca3a20e8eba43c8ff9d672793618`

## Changes since v1.2.3

### Other notable changes

* Ensure status is not changed during an update of PV, PVC, HPA objects ([#24924](https://github.com/kubernetes/kubernetes/pull/24924), [@mqliang](https://github.com/mqliang))
* GCI: Add two GCI specific metadata pairs ([#25105](https://github.com/kubernetes/kubernetes/pull/25105), [@andyzheng0831](https://github.com/andyzheng0831))
* Update salt config to allow Debian Jessie on GCE. ([#25123](https://github.com/kubernetes/kubernetes/pull/25123), [@jlewi](https://github.com/jlewi))
* Fix DeletingLoadBalancer event generation. ([#24833](https://github.com/kubernetes/kubernetes/pull/24833), [@a-robinson](https://github.com/a-robinson))
* GCE: Prefer preconfigured node tags for firewalls, if available ([#25148](https://github.com/kubernetes/kubernetes/pull/25148), [@a-robinson](https://github.com/a-robinson))
* Drain pods created from ReplicaSets in 'kubectl drain' ([#23689](https://github.com/kubernetes/kubernetes/pull/23689), [@maclof](https://github.com/maclof))
* GCI: Update the command to get the image ([#24987](https://github.com/kubernetes/kubernetes/pull/24987), [@andyzheng0831](https://github.com/andyzheng0831))
* Validate deletion timestamp doesn't change on update ([#24839](https://github.com/kubernetes/kubernetes/pull/24839), [@liggitt](https://github.com/liggitt))
* Add support for running clusters on GCI ([#24893](https://github.com/kubernetes/kubernetes/pull/24893), [@andyzheng0831](https://github.com/andyzheng0831))
* Trusty: Add retry in curl commands ([#24749](https://github.com/kubernetes/kubernetes/pull/24749), [@andyzheng0831](https://github.com/andyzheng0831))


# v1.3.0-alpha.3

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/master/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.3.0-alpha.3/kubernetes.tar.gz) | `01e0dc68653173614dc99f44875173478f837b38` | `ae22c35f3a963743d21daa17683e0288`

## Changes since v1.3.0-alpha.2

### Action Required

* Updating go-restful to generate "type":"object" instead of "type":"any" in swagger-spec (breaks kubectl 1.1) ([#22897](https://github.com/kubernetes/kubernetes/pull/22897), [@nikhiljindal](https://github.com/nikhiljindal))
* Make watch cache treat resourceVersion consistent with uncached watch ([#24008](https://github.com/kubernetes/kubernetes/pull/24008), [@liggitt](https://github.com/liggitt))

### Other notable changes

* Trusty: Add retry in curl commands ([#24749](https://github.com/kubernetes/kubernetes/pull/24749), [@andyzheng0831](https://github.com/andyzheng0831))
* Collect and expose runtime's image storage usage via Kubelet's /stats/summary endpoint ([#23595](https://github.com/kubernetes/kubernetes/pull/23595), [@vishh](https://github.com/vishh))
* Adding loadBalancer services to quota system ([#24247](https://github.com/kubernetes/kubernetes/pull/24247), [@sdminonne](https://github.com/sdminonne))
* Enforce --max-pods in kubelet admission; previously was only enforced in scheduler ([#24674](https://github.com/kubernetes/kubernetes/pull/24674), [@gmarek](https://github.com/gmarek))
* All clients under ClientSet share one RateLimiter. ([#24166](https://github.com/kubernetes/kubernetes/pull/24166), [@gmarek](https://github.com/gmarek))
* Remove requirement that Endpoints IPs be IPv4 ([#23317](https://github.com/kubernetes/kubernetes/pull/23317), [@aanm](https://github.com/aanm))
* Fix unintended change of Service.spec.ports[].nodePort during kubectl apply ([#24180](https://github.com/kubernetes/kubernetes/pull/24180), [@AdoHe](https://github.com/AdoHe))
* Don't log private SSH key ([#24506](https://github.com/kubernetes/kubernetes/pull/24506), [@timstclair](https://github.com/timstclair))
* Incremental improvements to kubelet e2e tests ([#24426](https://github.com/kubernetes/kubernetes/pull/24426), [@pwittrock](https://github.com/pwittrock))
* Bridge off-cluster traffic into services by masquerading. ([#24429](https://github.com/kubernetes/kubernetes/pull/24429), [@cjcullen](https://github.com/cjcullen))
* Flush conntrack state for removed/changed UDP Services ([#22573](https://github.com/kubernetes/kubernetes/pull/22573), [@freehan](https://github.com/freehan))
* Allow setting the Host header in a httpGet probe ([#24292](https://github.com/kubernetes/kubernetes/pull/24292), [@errm](https://github.com/errm))
* Fix goroutine leak in ssh-tunnel healthcheck. ([#24487](https://github.com/kubernetes/kubernetes/pull/24487), [@cjcullen](https://github.com/cjcullen))
* Fix gce.getDiskByNameUnknownZone logic. ([#24452](https://github.com/kubernetes/kubernetes/pull/24452), [@a-robinson](https://github.com/a-robinson))
* Make etcd cache size configurable ([#23914](https://github.com/kubernetes/kubernetes/pull/23914), [@jsravn](https://github.com/jsravn))
* Drain pods created from ReplicaSets ([#23689](https://github.com/kubernetes/kubernetes/pull/23689), [@maclof](https://github.com/maclof))
* Make kubectl edit not convert GV on edits ([#23437](https://github.com/kubernetes/kubernetes/pull/23437), [@DirectXMan12](https://github.com/DirectXMan12))
* don't ship kube-registry-proxy and pause images in tars. ([#23605](https://github.com/kubernetes/kubernetes/pull/23605), [@mikedanese](https://github.com/mikedanese))
* Do not throw creation errors for containers that fail immediately after being started ([#23894](https://github.com/kubernetes/kubernetes/pull/23894), [@vishh](https://github.com/vishh))
* Add a client flag to delete "--now" for grace period 0 ([#23756](https://github.com/kubernetes/kubernetes/pull/23756), [@smarterclayton](https://github.com/smarterclayton))
* add act-as powers ([#23549](https://github.com/kubernetes/kubernetes/pull/23549), [@deads2k](https://github.com/deads2k))
* Build Kubernetes, etcd and flannel for arm64 and ppc64le ([#23931](https://github.com/kubernetes/kubernetes/pull/23931), [@luxas](https://github.com/luxas))
* Honor starting resourceVersion in watch cache ([#24208](https://github.com/kubernetes/kubernetes/pull/24208), [@ncdc](https://github.com/ncdc))
* Update the pause image to build for arm64 and ppc64le ([#23697](https://github.com/kubernetes/kubernetes/pull/23697), [@luxas](https://github.com/luxas))
* Return more useful error information when a persistent volume fails to mount ([#23122](https://github.com/kubernetes/kubernetes/pull/23122), [@screeley44](https://github.com/screeley44))
* Trusty: Avoid unnecessary in-memory temp files ([#24144](https://github.com/kubernetes/kubernetes/pull/24144), [@andyzheng0831](https://github.com/andyzheng0831))
* e2e: fix error checking in kubelet stats ([#24205](https://github.com/kubernetes/kubernetes/pull/24205), [@yujuhong](https://github.com/yujuhong))
* Fixed mounting with containerized kubelet ([#23435](https://github.com/kubernetes/kubernetes/pull/23435), [@jsafrane](https://github.com/jsafrane))
* Adding nodeports services to quota ([#22154](https://github.com/kubernetes/kubernetes/pull/22154), [@sdminonne](https://github.com/sdminonne))
* e2e: adapt kubelet_perf.go to use the new summary metrics API ([#24003](https://github.com/kubernetes/kubernetes/pull/24003), [@yujuhong](https://github.com/yujuhong))
* kubelet: add RSS memory to the summary API ([#24015](https://github.com/kubernetes/kubernetes/pull/24015), [@yujuhong](https://github.com/yujuhong))

# v1.2.3

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/release-1.2/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.2.3/kubernetes.tar.gz) | `b2ce4e0c72562d09ba06e3c0913f0bd78da0285e` | `69e75650de30d5a52d144799e94a168d`

## Changes since v1.2.2

### Action Required

* Make watch cache treat resourceVersion consistent with uncached watch ([#24008](https://github.com/kubernetes/kubernetes/pull/24008), [@liggitt](https://github.com/liggitt))

### Other notable changes

* Fix unintended change of Service.spec.ports[].nodePort during kubectl apply ([#24180](https://github.com/kubernetes/kubernetes/pull/24180), [@AdoHe](https://github.com/AdoHe))
* Flush conntrack state for removed/changed UDP Services ([#22573](https://github.com/kubernetes/kubernetes/pull/22573), [@freehan](https://github.com/freehan))
* Allow setting the Host header in a httpGet probe ([#24292](https://github.com/kubernetes/kubernetes/pull/24292), [@errm](https://github.com/errm))
* Bridge off-cluster traffic into services by masquerading. ([#24429](https://github.com/kubernetes/kubernetes/pull/24429), [@cjcullen](https://github.com/cjcullen))
* Version-guard Kubectl client Guestbook application test against deployments ([#24478](https://github.com/kubernetes/kubernetes/pull/24478), [@ihmccreery](https://github.com/ihmccreery))
* Fix goroutine leak in ssh-tunnel healthcheck. ([#24487](https://github.com/kubernetes/kubernetes/pull/24487), [@cjcullen](https://github.com/cjcullen))
* Fixed mounting with containerized kubelet ([#23435](https://github.com/kubernetes/kubernetes/pull/23435), [@jsafrane](https://github.com/jsafrane))
* Do not throw creation errors for containers that fail immediately after being started ([#23894](https://github.com/kubernetes/kubernetes/pull/23894), [@vishh](https://github.com/vishh))
* Honor starting resourceVersion in watch cache ([#24208](https://github.com/kubernetes/kubernetes/pull/24208), [@ncdc](https://github.com/ncdc))
* Fix TerminationMessagePath ([#23658](https://github.com/kubernetes/kubernetes/pull/23658), [@Random-Liu](https://github.com/Random-Liu))
* Fix gce.getDiskByNameUnknownZone logic. ([#24452](https://github.com/kubernetes/kubernetes/pull/24452), [@a-robinson](https://github.com/a-robinson))
* kubelet: add RSS memory to the summary API ([#24015](https://github.com/kubernetes/kubernetes/pull/24015), [@yujuhong](https://github.com/yujuhong))
* e2e: adapt kubelet_perf.go to use the new summary metrics API ([#24003](https://github.com/kubernetes/kubernetes/pull/24003), [@yujuhong](https://github.com/yujuhong))
* e2e: fix error checking in kubelet stats ([#24205](https://github.com/kubernetes/kubernetes/pull/24205), [@yujuhong](https://github.com/yujuhong))
* Trusty: Avoid unnecessary in-memory temp files ([#24144](https://github.com/kubernetes/kubernetes/pull/24144), [@andyzheng0831](https://github.com/andyzheng0831))
* Allowing type object in kubectl swagger validation ([#24054](https://github.com/kubernetes/kubernetes/pull/24054), [@nikhiljindal](https://github.com/nikhiljindal))
* Add ClusterUpgrade tests ([#24150](https://github.com/kubernetes/kubernetes/pull/24150), [@ihmccreery](https://github.com/ihmccreery))
* Trusty: Do not create the docker-daemon cgroup ([#23996](https://github.com/kubernetes/kubernetes/pull/23996), [@andyzheng0831](https://github.com/andyzheng0831))


# v1.3.0-alpha.2

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/master/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.3.0-alpha.2/kubernetes.tar.gz) | `305c8c2af7e99d463dbbe4208ecfe2b50585e796` | `aadb8d729d855e69212008f8fda628c0`

## Changes since v1.3.0-alpha.1

### Other notable changes

* Make kube2sky and skydns docker images cross-platform ([#19376](https://github.com/kubernetes/kubernetes/pull/19376), [@luxas](https://github.com/luxas))
* Allowing type object in kubectl swagger validation ([#24054](https://github.com/kubernetes/kubernetes/pull/24054), [@nikhiljindal](https://github.com/nikhiljindal))
* Trusty: Do not create the docker-daemon cgroup ([#23996](https://github.com/kubernetes/kubernetes/pull/23996), [@andyzheng0831](https://github.com/andyzheng0831))
* Make ConfigMap volume readable as non-root ([#23793](https://github.com/kubernetes/kubernetes/pull/23793), [@pmorie](https://github.com/pmorie))
* only include running and pending pods in daemonset should place calculation ([#23929](https://github.com/kubernetes/kubernetes/pull/23929), [@mikedanese](https://github.com/mikedanese))
* Upgrade to golang 1.6 ([#22149](https://github.com/kubernetes/kubernetes/pull/22149), [@luxas](https://github.com/luxas))
* Cross-build hyperkube and debian-iptables for ARM. Also add a flannel image ([#21617](https://github.com/kubernetes/kubernetes/pull/21617), [@luxas](https://github.com/luxas))
* Add a timeout to the sshDialer to prevent indefinite hangs. ([#23843](https://github.com/kubernetes/kubernetes/pull/23843), [@cjcullen](https://github.com/cjcullen))
* Ensure object returned by volume getCloudProvider incorporates cloud config ([#23769](https://github.com/kubernetes/kubernetes/pull/23769), [@saad-ali](https://github.com/saad-ali))
* Update Dashboard UI addon to v1.0.1 ([#23724](https://github.com/kubernetes/kubernetes/pull/23724), [@maciaszczykm](https://github.com/maciaszczykm))
* Add zsh completion for kubectl ([#23797](https://github.com/kubernetes/kubernetes/pull/23797), [@sttts](https://github.com/sttts))
* AWS kube-up: tolerate a lack of ephemeral volumes ([#23776](https://github.com/kubernetes/kubernetes/pull/23776), [@justinsb](https://github.com/justinsb))
* duplicate kube-apiserver to federated-apiserver ([#23509](https://github.com/kubernetes/kubernetes/pull/23509), [@jianhuiz](https://github.com/jianhuiz))
* Kubelet: Start using the official docker engine-api ([#23506](https://github.com/kubernetes/kubernetes/pull/23506), [@Random-Liu](https://github.com/Random-Liu))
* Fix so setup-files don't recreate/invalidate certificates that already exist ([#23550](https://github.com/kubernetes/kubernetes/pull/23550), [@luxas](https://github.com/luxas))
* A pod never terminated if a container image registry was unavailable ([#23746](https://github.com/kubernetes/kubernetes/pull/23746), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Fix jsonpath to handle maps with key of nonstring types ([#23358](https://github.com/kubernetes/kubernetes/pull/23358), [@aveshagarwal](https://github.com/aveshagarwal))
* Trusty: Regional release .tar.gz support ([#23558](https://github.com/kubernetes/kubernetes/pull/23558), [@andyzheng0831](https://github.com/andyzheng0831))
* Add support for 3rd party objects to kubectl ([#18835](https://github.com/kubernetes/kubernetes/pull/18835), [@brendandburns](https://github.com/brendandburns))
* Remove unnecessary override of /etc/init.d/docker on containervm image. ([#23593](https://github.com/kubernetes/kubernetes/pull/23593), [@dchen1107](https://github.com/dchen1107))
* make docker-checker more robust ([#23662](https://github.com/kubernetes/kubernetes/pull/23662), [@ArtfulCoder](https://github.com/ArtfulCoder))
* Change kube-proxy & fluentd CPU request to 20m/80m. ([#23646](https://github.com/kubernetes/kubernetes/pull/23646), [@cjcullen](https://github.com/cjcullen))
* Create a new Deployment in kube-system for every version. ([#23512](https://github.com/kubernetes/kubernetes/pull/23512), [@Q-Lee](https://github.com/Q-Lee))
* IngressTLS: allow secretName to be blank for SNI routing ([#23500](https://github.com/kubernetes/kubernetes/pull/23500), [@tam7t](https://github.com/tam7t))
* don't sync deployment when pod selector is empty ([#23467](https://github.com/kubernetes/kubernetes/pull/23467), [@mikedanese](https://github.com/mikedanese))
* AWS: Fix problems with >2 security groups ([#23340](https://github.com/kubernetes/kubernetes/pull/23340), [@justinsb](https://github.com/justinsb))


# v1.2.2

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/release-1.2/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.2.2/kubernetes.tar.gz) | `8dede5833a1986434adea80749624f81a0db7bb4` | `72a5389f22827fb5133fdc3b7bfb9b3a`

## Changes since v1.2.1

### Other notable changes

* Trusty: Update heapster manifest handling code ([#23434](https://github.com/kubernetes/kubernetes/pull/23434), [@andyzheng0831](https://github.com/andyzheng0831))
* Support addon Deployments, make heapster a deployment with a nanny. ([#22893](https://github.com/kubernetes/kubernetes/pull/22893), [@Q-Lee](https://github.com/Q-Lee))
* Create a new Deployment in kube-system for every version. ([#23512](https://github.com/kubernetes/kubernetes/pull/23512), [@Q-Lee](https://github.com/Q-Lee))
* Use SCP to dump logs and parallelize a bit. ([#22835](https://github.com/kubernetes/kubernetes/pull/22835), [@spxtr](https://github.com/spxtr))
* Trusty: Regional release .tar.gz support ([#23558](https://github.com/kubernetes/kubernetes/pull/23558), [@andyzheng0831](https://github.com/andyzheng0831))
* Make ConfigMap volume readable as non-root ([#23793](https://github.com/kubernetes/kubernetes/pull/23793), [@pmorie](https://github.com/pmorie))
* only include running and pending pods in daemonset should place calculation ([#23929](https://github.com/kubernetes/kubernetes/pull/23929), [@mikedanese](https://github.com/mikedanese))
* A pod never terminated if a container image registry was unavailable ([#23746](https://github.com/kubernetes/kubernetes/pull/23746), [@derekwaynecarr](https://github.com/derekwaynecarr))
* Update Dashboard UI addon to v1.0.1 ([#23724](https://github.com/kubernetes/kubernetes/pull/23724), [@maciaszczykm](https://github.com/maciaszczykm))
* Ensure object returned by volume getCloudProvider incorporates cloud config ([#23769](https://github.com/kubernetes/kubernetes/pull/23769), [@saad-ali](https://github.com/saad-ali))
* Add a timeout to the sshDialer to prevent indefinite hangs. ([#23843](https://github.com/kubernetes/kubernetes/pull/23843), [@cjcullen](https://github.com/cjcullen))
* AWS kube-up: tolerate a lack of ephemeral volumes ([#23776](https://github.com/kubernetes/kubernetes/pull/23776), [@justinsb](https://github.com/justinsb))
* Fix so setup-files don't recreate/invalidate certificates that already exist ([#23550](https://github.com/kubernetes/kubernetes/pull/23550), [@luxas](https://github.com/luxas))

# v1.2.1

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/release-1.2/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.2.1/kubernetes.tar.gz) | `1639807c5788e1c6b1ab51fd30b723fb5debd865` | `235a1da47972c96a560d718d3256ca4f`

## Changes since v1.2.0

### Other notable changes

* AWS: Fix problems with >2 security groups ([#23340](https://github.com/kubernetes/kubernetes/pull/23340), [@justinsb](https://github.com/justinsb))
* IngressTLS: allow secretName to be blank for SNI routing ([#23500](https://github.com/kubernetes/kubernetes/pull/23500), [@tam7t](https://github.com/tam7t))
* Heapster patch release to 1.0.2 ([#23487](https://github.com/kubernetes/kubernetes/pull/23487), [@piosz](https://github.com/piosz))
* Remove unnecessary override of /etc/init.d/docker on containervm image. ([#23593](https://github.com/kubernetes/kubernetes/pull/23593), [@dchen1107](https://github.com/dchen1107))
* Change kube-proxy & fluentd CPU request to 20m/80m. ([#23646](https://github.com/kubernetes/kubernetes/pull/23646), [@cjcullen](https://github.com/cjcullen))
* make docker-checker more robust ([#23662](https://github.com/kubernetes/kubernetes/pull/23662), [@ArtfulCoder](https://github.com/ArtfulCoder))
* validate that daemonsets don't have empty selectors on creation ([#23530](https://github.com/kubernetes/kubernetes/pull/23530), [@mikedanese](https://github.com/mikedanese))
* don't sync deployment when pod selector is empty ([#23467](https://github.com/kubernetes/kubernetes/pull/23467), [@mikedanese](https://github.com/mikedanese))
* Support differentiation of OS distro in e2e tests ([#23466](https://github.com/kubernetes/kubernetes/pull/23466), [@andyzheng0831](https://github.com/andyzheng0831))
* don't sync daemonsets with selectors that match all pods ([#23223](https://github.com/kubernetes/kubernetes/pull/23223), [@mikedanese](https://github.com/mikedanese))
* Trusty: Avoid reaching GCE custom metadata size limit ([#22818](https://github.com/kubernetes/kubernetes/pull/22818), [@andyzheng0831](https://github.com/andyzheng0831))
* Update kubectl help for 1.2 resources ([#23305](https://github.com/kubernetes/kubernetes/pull/23305), [@janetkuo](https://github.com/janetkuo))
* Removing URL query param from swagger UI to fix the XSS issue ([#23234](https://github.com/kubernetes/kubernetes/pull/23234), [@nikhiljindal](https://github.com/nikhiljindal))
* Fix hairpin mode ([#23325](https://github.com/kubernetes/kubernetes/pull/23325), [@MurgaNikolay](https://github.com/MurgaNikolay))
* Bump to container-vm-v20160321 ([#23313](https://github.com/kubernetes/kubernetes/pull/23313), [@zmerlynn](https://github.com/zmerlynn))
* Remove the restart-kube-proxy and restart-apiserver functions ([#23180](https://github.com/kubernetes/kubernetes/pull/23180), [@roberthbailey](https://github.com/roberthbailey))
* Copy annotations back from RS to Deployment on rollback ([#23160](https://github.com/kubernetes/kubernetes/pull/23160), [@janetkuo](https://github.com/janetkuo))
* Trusty: Support hybrid cluster with nodes on ContainerVM ([#23079](https://github.com/kubernetes/kubernetes/pull/23079), [@andyzheng0831](https://github.com/andyzheng0831))
* update expose command description to add deployment ([#23246](https://github.com/kubernetes/kubernetes/pull/23246), [@AdoHe](https://github.com/AdoHe))
* Add a rate limiter to the GCE cloudprovider ([#23019](https://github.com/kubernetes/kubernetes/pull/23019), [@alex-mohr](https://github.com/alex-mohr))
* Add a Deployment example for kubectl expose. ([#23222](https://github.com/kubernetes/kubernetes/pull/23222), [@madhusudancs](https://github.com/madhusudancs))
* Use versioned object when computing patch ([#23145](https://github.com/kubernetes/kubernetes/pull/23145), [@liggitt](https://github.com/liggitt))
* kubelet: send all recevied pods in one update ([#23141](https://github.com/kubernetes/kubernetes/pull/23141), [@yujuhong](https://github.com/yujuhong))
* Add a SSHKey sync check to the master's healthz (when using SSHTunnels). ([#23167](https://github.com/kubernetes/kubernetes/pull/23167), [@cjcullen](https://github.com/cjcullen))
* Validate minimum CPU limits to be >= 10m ([#23143](https://github.com/kubernetes/kubernetes/pull/23143), [@vishh](https://github.com/vishh))
* kubernetes/kubernetes#23034 Fix controller-manager race condition issue which cause endpoints flush during restart ([#23035](https://github.com/kubernetes/kubernetes/pull/23035), [@xinxiaogang](https://github.com/xinxiaogang))
* MESOS: forward globally declared cadvisor housekeeping flags ([#22974](https://github.com/kubernetes/kubernetes/pull/22974), [@jdef](https://github.com/jdef))
* Trusty: support developer workflow on base image ([#22960](https://github.com/kubernetes/kubernetes/pull/22960), [@andyzheng0831](https://github.com/andyzheng0831))


# v1.3.0-alpha.1

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/HEAD/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.3.0-alpha.1/kubernetes.tar.gz) | `e0041b08e220a4704ea2ad90a6ec7c8f2120c2d3` | `7bb2df32aea94678f72a8d1f43a12098`

## Changes since v1.2.0

### Action Required

* Disabling swagger ui by default on apiserver. Adding a flag that can enable it ([#23025](https://github.com/kubernetes/kubernetes/pull/23025), [@nikhiljindal](https://github.com/nikhiljindal))
* restore ability to run against secured etcd ([#21535](https://github.com/kubernetes/kubernetes/pull/21535), [@AdoHe](https://github.com/AdoHe))

### Other notable changes

* Trusty: Avoid reaching GCE custom metadata size limit ([#22818](https://github.com/kubernetes/kubernetes/pull/22818), [@andyzheng0831](https://github.com/andyzheng0831))
* Update kubectl help for 1.2 resources ([#23305](https://github.com/kubernetes/kubernetes/pull/23305), [@janetkuo](https://github.com/janetkuo))
* Removing URL query param from swagger UI to fix the XSS issue ([#23234](https://github.com/kubernetes/kubernetes/pull/23234), [@nikhiljindal](https://github.com/nikhiljindal))
* Fix hairpin mode ([#23325](https://github.com/kubernetes/kubernetes/pull/23325), [@MurgaNikolay](https://github.com/MurgaNikolay))
* Bump to container-vm-v20160321 ([#23313](https://github.com/kubernetes/kubernetes/pull/23313), [@zmerlynn](https://github.com/zmerlynn))
* Remove the restart-kube-proxy and restart-apiserver functions ([#23180](https://github.com/kubernetes/kubernetes/pull/23180), [@roberthbailey](https://github.com/roberthbailey))
* Copy annotations back from RS to Deployment on rollback ([#23160](https://github.com/kubernetes/kubernetes/pull/23160), [@janetkuo](https://github.com/janetkuo))
* Trusty: Support hybrid cluster with nodes on ContainerVM ([#23079](https://github.com/kubernetes/kubernetes/pull/23079), [@andyzheng0831](https://github.com/andyzheng0831))
* update expose command description to add deployment ([#23246](https://github.com/kubernetes/kubernetes/pull/23246), [@AdoHe](https://github.com/AdoHe))
* Add a rate limiter to the GCE cloudprovider ([#23019](https://github.com/kubernetes/kubernetes/pull/23019), [@alex-mohr](https://github.com/alex-mohr))
* Add a Deployment example for kubectl expose. ([#23222](https://github.com/kubernetes/kubernetes/pull/23222), [@madhusudancs](https://github.com/madhusudancs))
* Use versioned object when computing patch ([#23145](https://github.com/kubernetes/kubernetes/pull/23145), [@liggitt](https://github.com/liggitt))
* kubelet: send all recevied pods in one update ([#23141](https://github.com/kubernetes/kubernetes/pull/23141), [@yujuhong](https://github.com/yujuhong))
* Add a SSHKey sync check to the master's healthz (when using SSHTunnels). ([#23167](https://github.com/kubernetes/kubernetes/pull/23167), [@cjcullen](https://github.com/cjcullen))
* Validate minimum CPU limits to be >= 10m ([#23143](https://github.com/kubernetes/kubernetes/pull/23143), [@vishh](https://github.com/vishh))
* kubernetes/kubernetes#23034 Fix controller-manager race condition issue which cause endpoints flush during restart ([#23035](https://github.com/kubernetes/kubernetes/pull/23035), [@xinxiaogang](https://github.com/xinxiaogang))
* MESOS: forward globally declared cadvisor housekeeping flags ([#22974](https://github.com/kubernetes/kubernetes/pull/22974), [@jdef](https://github.com/jdef))
* Trusty: support developer workflow on base image ([#22960](https://github.com/kubernetes/kubernetes/pull/22960), [@andyzheng0831](https://github.com/andyzheng0831))
* Bumped Heapster to stable version 1.0.0 ([#22993](https://github.com/kubernetes/kubernetes/pull/22993), [@piosz](https://github.com/piosz))
* Deprecating --api-version flag ([#22410](https://github.com/kubernetes/kubernetes/pull/22410), [@nikhiljindal](https://github.com/nikhiljindal))
* allow resource.version.group in kubectl ([#22853](https://github.com/kubernetes/kubernetes/pull/22853), [@deads2k](https://github.com/deads2k))
* Use SCP to dump logs and parallelize a bit. ([#22835](https://github.com/kubernetes/kubernetes/pull/22835), [@spxtr](https://github.com/spxtr))
* update wide option output ([#22772](https://github.com/kubernetes/kubernetes/pull/22772), [@AdoHe](https://github.com/AdoHe))
* Change scheduler logic from random to round-robin ([#22430](https://github.com/kubernetes/kubernetes/pull/22430), [@gmarek](https://github.com/gmarek))

# v1.2.0

[Documentation](http://kubernetes.github.io) & [Examples](http://releases.k8s.io/release-1.2/examples)

## Downloads

binary | sha1 hash | md5 hash
------ | --------- | --------
[kubernetes.tar.gz](https://storage.googleapis.com/kubernetes-release/release/v1.2.0/kubernetes.tar.gz) | `52dd998e1191f464f581a9b87017d70ce0b058d9` | `c0ce9e6150e9d7a19455db82f3318b4c`

## Changes since v1.1.1

### Major Themes

  * <strong>Significant scale improvements</strong>. Increased cluster scale by 400% to 1000 nodes with 30,000 pods per cluster.
Kubelet supports 100 pods per node with 4x reduced system overhead.
  * <strong>Simplified application deployment and management. </strong>
     * Dynamic Configuration (ConfigMap API in the core API group) enables application
configuration to be stored as a Kubernetes API object and pulled dynamically on
container startup, as an alternative to baking in command-line flags when a
container is built.
     * Turnkey Deployments (Deployment API (Beta) in the Extensions API group)
automate deployment and rolling updates of applications, specified
declaratively. It handles versioning, multiple simultaneous rollouts,
aggregating status across all pods, maintaining application availability, and
rollback.
  * <strong>Automated cluster management: </strong>
     * Kubernetes clusters can now span zones within a cloud provider. Pods from a
service will be automatically spread across zones, enabling applications to
tolerate zone failure.
     * Simplified way to run a container on every node (DaemonSet API (Beta) in the
Extensions API group): Kubernetes can schedule a service (such as a logging
agent) that runs one, and only one, pod per node.
     * TLS and L7 support (Ingress API (Beta) in the Extensions API group): Kubernetes
is now easier to integrate into custom networking environments by supporting
TLS for secure communication and L7 http-based traffic routing.
     * Graceful Node Shutdown (aka drain) - The new “kubectl drain” command gracefully
evicts pods from nodes in preparation for disruptive operations like kernel
upgrades or maintenance.
     * Custom Metrics for Autoscaling (HorizontalPodAutoscaler API in the Autoscaling
API group): The Horizontal Pod Autoscaling feature now supports custom metrics
(Alpha), allowing you to specify application-level metrics and thresholds to
trigger scaling up and down the number of pods in your application.
  * <strong>New GUI</strong> (dashboard) allows you to get started quickly and enables the same
functionality found in the CLI as a more approachable and discoverable way of
interacting with the system. Note: the GUI is enabled by default in 1.2 clusters.

<img src="docs/images/newgui.png" width="" alt="Dashboard UI screenshot showing cards that represent applications that run inside a cluster" title="Dashboard UI apps screen">

### Other notable improvements

  * Job was Beta in 1.1 and is GA in 1.2 .
     * <code>apiVersion: batch/v1 </code>is now available.  You now do not need to specify the <code>.spec.selector</code> field — a [unique selector is automatically generated ](http://kubernetes.io/docs/user-guide/jobs/#pod-selector)for you.
     * The previous version, <code>apiVersion: extensions/v1beta1</code>, is still supported.  Even if you roll back to 1.1, the objects created using
the new apiVersion will still be accessible, using the old version.   You can
continue to use your existing JSON and YAML files until you are ready to switch
to <code>batch/v1</code>.  We may remove support for Jobs with  <code>apiVersion: extensions/v1beta1 </code>in 1.3 or 1.4.
  *  HorizontalPodAutoscaler was Beta in 1.1 and is GA in 1.2 .
     * <code>apiVersion: autoscaling/v1 </code>is now available.  Changes in this version are:
        * Field CPUUtilization which was a nested structure CPUTargetUtilization in
HorizontalPodAutoscalerSpec was replaced by TargetCPUUtilizationPercentage
which is an integer.
        * ScaleRef of type SubresourceReference in HorizontalPodAutoscalerSpec which
referred to scale subresource of the resource being scaled was replaced by
ScaleTargetRef which points just to the resource being scaled.
        * In extensions/v1beta1 if CPUUtilization in HorizontalPodAutoscalerSpec was not
specified it was set to 80 by default while in autoscaling/v1 HPA object
without TargetCPUUtilizationPercentage specified is a valid object. Pod
autoscaler controller will apply a default scaling policy in this case which is
equivalent to the previous one but may change in the future.
     * The previous version, <code>apiVersion: extensions/v1beta1</code>, is still supported.  Even if you roll back to 1.1, the objects created using
the new apiVersions will still be accessible, using the old version.  You can
continue to use your existing JSON and YAML files until you are ready to switch
to <code>autoscaling/v1</code>.  We may remove support for HorizontalPodAutoscalers with  <code>apiVersion: extensions/v1beta1 </code>in 1.3 or 1.4.
  * Kube-Proxy now defaults to an iptables-based proxy. If the --proxy-mode flag is
specified while starting kube-proxy (‘userspace’ or ‘iptables’), the flag value
will be respected. If the flag value is not specified, the kube-proxy respects
the Node object annotation: ‘net.beta.kubernetes.io/proxy-mode’. If the
annotation is not specified, then ‘iptables’ mode is the default. If kube-proxy
is unable to start in iptables mode because system requirements are not met
(kernel or iptables versions are insufficient), the kube-proxy will fall-back
to userspace mode. Kube-proxy is much more performant and less
resource-intensive in ‘iptables’ mode.
  * Node stability can be improved by reserving [resources](https://github.com/kubernetes/kubernetes/blob/release-1.2/docs/proposals/node-allocatable.md) for the base operating system using --system-reserved and --kube-reserved Kubelet flags
  * Liveness and readiness probes now support more configuration parameters:
periodSeconds, successThreshold, failureThreshold
  * The new ReplicaSet API (Beta) in the Extensions API group is similar to
ReplicationController, but its [selector](http://kubernetes.io/docs/user-guide/labels/#label-selectors) is more general (supports set-based selector; whereas ReplicationController
only supports equality-based selector).
  * Scale subresource support is now expanded to ReplicaSets along with
ReplicationControllers and Deployments. Scale now supports two different types
of selectors to accommodate both [equality-based selectors](http://kubernetes.io/docs/user-guide/labels/#equality-based-requirement) supported by ReplicationControllers and [set-based selectors](http://kubernetes.io/docs/user-guide/labels/#set-based-requirement) supported by Deployments and ReplicaSets.
  * “kubectl run” now produces Deployments (instead of ReplicationControllers) and
Jobs (instead of Pods) by default.
  * Pods can now consume Secret data in environment variables and inject those
environment variables into a container’s command-line args.
  * Stable version of Heapster which scales up to 1000 nodes: more metrics, reduced
latency, reduced cpu/memory consumption (~4mb per monitored node).
  * Pods now have a security context which allows users to specify:
     * attributes which apply to the whole pod:
        * User ID
        * Whether all containers should be non-root
        * Supplemental Groups
        * FSGroup - a special supplemental group
        * SELinux options
     * If a pod defines an FSGroup, that Pod’s system (emptyDir, secret, configMap,
etc) volumes and block-device volumes will be owned by the FSGroup, and each
container in the pod will run with the FSGroup as a supplemental group
  * Volumes that support SELinux labelling are now automatically relabeled with the
Pod’s SELinux context, if specified
  * A stable client library release\_1\_2 is added. The library is [here](pkg/client/clientset_generated/release_1_2/), and detailed doc is [here](docs/devel/generating-clientset.md#released-clientsets). We will keep the interface of this go client stable.
  * New Azure File Service Volume Plugin enables mounting Microsoft Azure File
Volumes (SMB 2.1 and 3.0) into a Pod. See [example](https://github.com/kubernetes/kubernetes/blob/release-1.2/examples/azure_file/README.md) for details.
  * Logs usage and root filesystem usage of a container, volumes usage of a pod and node disk usage are exposed through Kubelet new metrics API.

### Experimental Features

  * Dynamic Provisioning of PersistentVolumes: Kubernetes previously required all
volumes to be manually provisioned by a cluster administrator before use. With
this feature, volume plugins that support it (GCE PD, AWS EBS, and Cinder) can
automatically provision a PersistentVolume to bind to an unfulfilled
PersistentVolumeClaim.
  * Run multiple schedulers in parallel, e.g. one or more custom schedulers
alongside the default Kubernetes scheduler, using pod annotations to select
among the schedulers for each pod. Documentation is [here](http://kubernetes.io/docs/admin/multiple-schedulers.md), design doc is [here](docs/proposals/multiple-schedulers.md).
  * More expressive node affinity syntax, and support for “soft” node affinity.
Node selectors (to constrain pods to schedule on a subset of nodes) now support
the operators {<code>In, NotIn, Exists, DoesNotExist, Gt, Lt</code>}  instead of just conjunction of exact match on node label values. In
addition, we’ve introduced a new “soft” kind of node selector that is just a
hint to the scheduler; the scheduler will try to satisfy these requests but it
does not guarantee they will be satisfied. Both the “hard” and “soft” variants
of node affinity use the new syntax. Documentation is [here](http://kubernetes.io/docs/user-guide/node-selection/) (see section “Alpha feature in Kubernetes v1.2: Node Affinity“). Design doc is [here](https://github.com/kubernetes/kubernetes/blob/release-1.2/docs/design/nodeaffinity.md).
  * A pod can specify its own Hostname and Subdomain via annotations (<code>pod.beta.kubernetes.io/hostname, pod.beta.kubernetes.io/subdomain)</code>. If the Subdomain matches the name of a [headless service](http://kubernetes.io/docs/user-guide/services/#headless-services) in the same namespace, a DNS A record is also created for the pod’s FQDN. More
details can be found in the [DNS README](https://github.com/kubernetes/kubernetes/blob/release-1.2/cluster/saltbase/salt/kube-dns/README.md#a-records-and-hostname-based-on-pod-annotations---a-beta-feature-in-kubernetes-v12). Changes were introduced in PR [#20688](https://github.com/kubernetes/kubernetes/pull/20688).
  * New SchedulerExtender enables users to implement custom
out-of-(the-scheduler)-process scheduling predicates and priority functions,
for example to schedule pods based on resources that are not directly managed
by Kubernetes. Changes were introduced in PR [#13580](https://github.com/kubernetes/kubernetes/pull/13580). Example configuration and documentation is available [here](docs/design/scheduler_extender.md). This is an alpha feature and may not be supported in its current form at beta
or GA.
  * New Flex Volume Plugin enables users to use out-of-process volume plugins that
are installed to “/usr/libexec/kubernetes/kubelet-plugins/volume/exec/” on
every node, instead of being compiled into the Kubernetes binary. See [example](examples/flexvolume/README.md) for details.
  * vendor volumes into a pod. It expects vendor drivers are installed in the
volume plugin path on each kubelet node. This is an alpha feature and may
change in future.
  * Kubelet exposes a new Alpha metrics API - /stats/summary in a user friendly format with reduced system overhead. The measurement is done in PR [#22542](https://github.com/kubernetes/kubernetes/pull/22542).

### Action required

  * Docker v1.9.1 is officially recommended. Docker v1.8.3 and Docker v1.10 are
supported. If you are using an older release of Docker, please upgrade. Known
issues with Docker 1.9.1 can be found below.
  * CPU hardcapping will be enabled by default for containers with CPU limit set,
if supported by the kernel. You should either adjust your CPU limit, or set CPU
request only, if you want to avoid hardcapping. If the kernel does not support
CPU Quota, NodeStatus will contain a warning indicating that CPU Limits cannot
be enforced.
  * The following applies only if you use the Go language client (<code>/pkg/client/unversioned</code>) to create Job by defining Go variables of type "<code>k8s.io/kubernetes/pkg/apis/extensions".Job</code>).  We think <strong>this is not common</strong>, so if you are not sure what this means, you probably aren't doing this.  If
you do this, then, at the time you re-vendor the "<code>k8s.io/kubernetes/"</code> code, you will need to set <code>job.Spec.ManualSelector = true</code>, or else set <code>job.Spec.Selector = nil.  </code>Otherwise, the jobs you create may be rejected.  See [Specifying your own pod selector](http://kubernetes.io/docs/user-guide/jobs/#specifying-your-own-pod-selector).
  * Deployment was Alpha in 1.1 (though it had apiVersion extensions/v1beta1) and
was disabled by default. Due to some non-backward-compatible API changes, any
Deployment objects you created in 1.1 won’t work with in the 1.2 release.
     * Before upgrading to 1.2, <strong>delete all Deployment alpha-version resources</strong>, including the Replication Controllers and Pods the Deployment manages. Then
create Deployment Beta resources after upgrading to 1.2. Not deleting the
Deployment objects may cause the deployment controller to mistakenly match
other pods and delete them, due to the selector API change.
     * Client (kubectl) and server versions must match (both 1.1 or both 1.2) for any
Deployment-related operations.
     * Behavior change:
        * Deployment creates ReplicaSets instead of ReplicationControllers.
        * Scale subresource now has a new <code>targetSelector</code> field in its status. This field supports the new set-based selectors supported
by Deployments, but in a serialized format.
     * Spec change:
        * Deployment’s [selector](http://kubernetes.io/docs/user-guide/labels/#label-selectors) is now more general (supports set-based selector; it only supported
equality-based selector in 1.1).
        * .spec.uniqueLabelKey is removed -- users can’t customize unique label key --
and its default value is changed from
“deployment.kubernetes.io/podTemplateHash” to “pod-template-hash”.
        * .spec.strategy.rollingUpdate.minReadySeconds is moved to .spec.minReadySeconds
  * DaemonSet was Alpha in 1.1 (though it had apiVersion extensions/v1beta1) and
was disabled by default. Due to some non-backward-compatible API changes, any
DaemonSet objects you created in 1.1 won’t work with in the 1.2 release.
     * Before upgrading to 1.2, <strong>delete all DaemonSet alpha-version resources</strong>. If you do not want to disrupt the pods, use kubectl delete daemonset <name>
--cascade=false. Then create DaemonSet Beta resources after upgrading to 1.2.
     * Client (kubectl) and server versions must match (both 1.1 or both 1.2) for any
DaemonSet-related operations.
     * Behavior change:
        * DaemonSet pods will be created on nodes with .spec.unschedulable=true and will
not be evicted from nodes whose Ready condition is false.
        * Updates to the pod template are now permitted. To perform a rolling update of a
DaemonSet, update the pod template and then delete its pods one by one; they
will be replaced using the updated template.
     * Spec change:
        * DaemonSet’s [selector](http://kubernetes.io/docs/user-guide/labels/#label-selectors) is now more general (supports set-based selector; it only supported
equality-based selector in 1.1).
  * Running against a secured etcd requires these flags to be passed to
kube-apiserver (instead of --etcd-config):
     * --etcd-certfile, --etcd-keyfile (if using client cert auth)
     * --etcd-cafile (if not using system roots)
  * As part of preparation in 1.2 for adding support for protocol buffers (and the
direct YAML support in the API available today), the Content-Type and Accept
headers are now properly handled as per the HTTP spec.  As a consequence, if
you had a client that was sending an invalid Content-Type or Accept header to
the API, in 1.2 you will either receive a 415 or 406 error.
The only client
this is known to affect is curl when you use -d with JSON but don't set a
content type, helpfully sends "application/x-www-urlencoded", which is not
correct.
Other client authors should double check that you are sending proper
accept and content type headers, or set no value (in which case JSON is the
default).
An example using curl:
<code>curl -H "Content-Type: application/json" -XPOST -d
'{"apiVersion":"v1","kind":"Namespace","metadata":{"name":"kube-system"}}' "[http://127.0.0.1:8080/api/v1/namespaces](http://127.0.0.1:8080/api/v1/namespaces)"</code>
  * The version of InfluxDB is bumped from 0.8 to 0.9 which means storage schema
change. More details [here](https://docs.influxdata.com/influxdb/v0.9/administration/upgrading/).
  * We have renamed “minions” to “nodes”.  If you were specifying NUM\_MINIONS or
MINION\_SIZE to kube-up, you should now specify NUM\_NODES or NODE\_SIZE.

### Known Issues

  * Paused deployments can't be resized and don't clean up old ReplicaSets.
  * Minimum memory limit is 4MB. This is a docker limitation
  * Minimum CPU limits is 10m. This is a Linux Kernel limitation
  * “kubectl rollout undo” (i.e. rollback) will hang on paused deployments, because
paused deployments can’t be rolled back (this is expected), and the command
waits for rollback events to return the result. Users should use “kubectl
rollout resume” to resume a deployment before rolling back.
  * “kubectl edit <list>” will open the editor multiple times, once for each
resource in the list.
  * If you create HPA object using autoscaling/v1 API without specifying
targetCPUUtilizationPercentage and read it using kubectl it will print default
value as specified in extensions/v1beta1 (see details in [#23196](https://github.com/kubernetes/kubernetes/issues/23196)).
  * If a node or kubelet crashes with a volume attached, the volume will remain
attached to that node. If that volume can only be attached to one node at a
time (GCE PDs attached in RW mode, for example), then the volume must be
manually detached before Kubernetes can attach it to other nodes.
  * If a volume is already attached to a node any subsequent attempts to attach it
again (due to kubelet restart, for example) will fail. The volume must either
be manually detached first or the pods referencing it deleted (which would
trigger automatic volume detach).
  * In very large clusters it may happen that a few nodes won’t register in API
server in a given timeframe for whatever reasons (networking issue, machine
failure, etc.). Normally when kube-up script will encounter even one NotReady
node it will fail, even though the cluster most likely will be working. We
added an environmental variable to kube-up ALLOWED\_NOTREADY\_NODES that
defines the number of nodes that if not Ready in time won’t cause kube-up
failure.
  * “kubectl rolling-update” only supports Replication Controllers (it doesn’t
support Replica Sets). It’s recommended to use Deployment 1.2 with “kubectl
rollout” commands instead, if you want to rolling update Replica Sets.
  * When live upgrading Kubelet to 1.2 without draining the pods running on the node,
the containers will be restarted by Kubelet (see details in [#23104](https://github.com/kubernetes/kubernetes/issues/23104)).

#### Docker Known Issues

##### 1.9.1

  * Listing containers can be slow at times which will affect kubelet performance.
More information [here](https://github.com/docker/docker/issues/17720)
  * Docker daemon restarts can fail. Docker checkpoints have to deleted between
restarts. More information [here](https://github.com/kubernetes/kubernetes/issues/20995)
  * Pod IP allocation-related issues. Deleting the docker checkpoint prior to
restarting the daemon alleviates this issue, but hasn’t been verified to
completely eliminate the IP allocation issue. More information [here](https://github.com/kubernetes/kubernetes/issues/21523#issuecomment-191498969)
  * Daemon becomes unresponsive (rarely) due to kernel deadlocks. More information [here](https://github.com/kubernetes/kubernetes/issues/21866#issuecomment-189492391)

### Provider-specific Notes

#### Various

   Core changes:

  * Support for load balancers with source ranges

#### AWS

Core changes:

  * Support for ELBs with complex configurations: better subnet selection with
multiple subnets, and internal ELBs
  * Support for VPCs with private dns names
  * Multiple fixes to EBS volume mounting code for robustness, and to support
mounting the full number of AWS recommended volumes.
  * Multiple fixes to avoid hitting AWS rate limits, and to throttle if we do
  * Support for the EC2 Container Registry (currently in us-east-1 only)

With kube-up:

  * Automatically install updates on boot & reboot
  * Use optimized image based on Jessie by default
  * Add support for Ubuntu Wily
  * Master is configured with automatic restart-on-failure, via CloudWatch
  * Bootstrap reworked to be more similar to GCE; better supports reboots/restarts
  * Use an elastic IP for the master by default
  * Experimental support for node spot instances (set NODE\_SPOT\_PRICE=0.05)

#### GCE

  * Ubuntu Trusty support added

Please see the [Releases Page](https://github.com/kubernetes/kubernetes/releases) for older releases.



[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/CHANGELOG.md?pixel)]()
