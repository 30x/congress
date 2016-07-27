#!/bin/bash

# Copyright 2015 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# A configuration for Kubemark cluster. It doesn't need to be kept in
# sync with gce/config-default.sh (except the filename, because I'm reusing
# gce/util.sh script which assumes config filename), but if some things that
# are enabled by default should not run in hollow clusters, they should be disabled here.

source "${KUBE_ROOT}/cluster/gce/config-common.sh"

GCLOUD=gcloud
ZONE=${KUBE_GCE_ZONE:-us-central1-b}
NUM_NODES=${NUM_NODES:-100}
MASTER_SIZE=${MASTER_SIZE:-n1-standard-$(get-master-size)}
MASTER_DISK_TYPE=pd-ssd
MASTER_DISK_SIZE=${MASTER_DISK_SIZE:-20GB}
REGISTER_MASTER_KUBELET=${REGISTER_MASTER:-false}
PREEMPTIBLE_NODE=${PREEMPTIBLE_NODE:-false}

OS_DISTRIBUTION=${KUBE_OS_DISTRIBUTION:-debian}
MASTER_IMAGE=${KUBE_GCE_MASTER_IMAGE:-container-vm-v20160321}
MASTER_IMAGE_PROJECT=${KUBE_GCE_MASTER_PROJECT:-google-containers}

NETWORK=${KUBE_GCE_NETWORK:-default}
INSTANCE_PREFIX="${INSTANCE_PREFIX:-"default"}"
MASTER_NAME="${INSTANCE_PREFIX}-kubemark-master"
MASTER_TAG="kubemark-master"
MASTER_IP_RANGE="${MASTER_IP_RANGE:-10.246.0.0/24}"
CLUSTER_IP_RANGE="${CLUSTER_IP_RANGE:-10.240.0.0/13}"
RUNTIME_CONFIG="${KUBE_RUNTIME_CONFIG:-}"
TERMINATED_POD_GC_THRESHOLD=${TERMINATED_POD_GC_THRESHOLD:-100}

# Default Log level for all components in test clusters and variables to override it in specific components.
TEST_CLUSTER_LOG_LEVEL="${TEST_CLUSTER_LOG_LEVEL:---v=2}"
KUBELET_TEST_LOG_LEVEL="${KUBELET_TEST_LOG_LEVEL:-$TEST_CLUSTER_LOG_LEVEL}"
API_SERVER_TEST_LOG_LEVEL="${API_SERVER_TEST_LOG_LEVEL:-$TEST_CLUSTER_LOG_LEVEL}"
CONTROLLER_MANAGER_TEST_LOG_LEVEL="${CONTROLLER_MANAGER_TEST_LOG_LEVEL:-$TEST_CLUSTER_LOG_LEVEL}"
SCHEDULER_TEST_LOG_LEVEL="${SCHEDULER_TEST_LOG_LEVEL:-$TEST_CLUSTER_LOG_LEVEL}"
KUBEPROXY_TEST_LOG_LEVEL="${KUBEPROXY_TEST_LOG_LEVEL:-$TEST_CLUSTER_LOG_LEVEL}"

TEST_CLUSTER_DELETE_COLLECTION_WORKERS="${TEST_CLUSTER_DELETE_COLLECTION_WORKERS:---delete-collection-workers=16}"
TEST_CLUSTER_RESYNC_PERIOD="${TEST_CLUSTER_RESYNC_PERIOD:-}"

# ContentType used by all components to communicate with apiserver.
TEST_CLUSTER_API_CONTENT_TYPE="${TEST_CLUSTER_API_CONTENT_TYPE:-}"
# ContentType used to store objects in underlying database.
TEST_CLUSTER_STORAGE_CONTENT_TYPE="${TEST_CLUSTER_STORAGE_CONTENT_TYPE:-}"

KUBELET_TEST_ARGS="--max-pods=100 $TEST_CLUSTER_LOG_LEVEL ${TEST_CLUSTER_API_CONTENT_TYPE}"
APISERVER_TEST_ARGS="--runtime-config=extensions/v1beta1 ${API_SERVER_TEST_LOG_LEVEL} ${TEST_CLUSTER_STORAGE_CONTENT_TYPE} ${TEST_CLUSTER_DELETE_COLLECTION_WORKERS}"
CONTROLLER_MANAGER_TEST_ARGS="${CONTROLLER_MANAGER_TEST_LOG_LEVEL} ${TEST_CLUSTER_RESYNC_PERIOD} ${TEST_CLUSTER_API_CONTENT_TYPE}"
SCHEDULER_TEST_ARGS="${SCHEDULER_TEST_LOG_LEVEL} ${TEST_CLUSTER_API_CONTENT_TYPE}"
KUBEPROXY_TEST_ARGS="${KUBEPROXY_TEST_LOG_LEVEL} ${TEST_CLUSTER_API_CONTENT_TYPE}"

SERVICE_CLUSTER_IP_RANGE="10.0.0.0/16"  # formerly PORTAL_NET
ALLOCATE_NODE_CIDRS=true
