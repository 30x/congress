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

package main

import (
  "log"

  "github.com/30x/congress/utils"
  "github.com/30x/congress/congress"
)

func main() {
  var doNamespaceRestart bool
  var doPolicyRestart bool

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
  namespaceWatcher, policyWatcher := congress.InitCongress(kubeClient, extClient, config)
  log.Println("Congress is in session")

  // begin reading namespace events
  namespaceCh := namespaceWatcher.ResultChan()
  policyCh := policyWatcher.ResultChan()
  for {
    select {
      case event, ok := <-namespaceCh:
        if !ok {
          log.Println("Kuberenetes closed the namespace watcher")
          doNamespaceRestart = true
        }

        if !doNamespaceRestart {
          // new namespace created event, start isolating it
          err := congress.HandleNamespaceEvent(event, kubeClient, extClient, config)
          if err != nil { log.Fatal(err) }
        }
      case event, ok := <-policyCh:
        if !ok {
          log.Println("Kubernetes closed the policy watcher")
          doPolicyRestart = true
        }

        if !doPolicyRestart {
          err := congress.HandlePolicyEvent(event, kubeClient, extClient, config)
          if err != nil { log.Fatal(err) }
        }
    }

    if doNamespaceRestart {
      log.Println("Congress is restarting the namespace watcher")
      doNamespaceRestart = false
      namespaceWatcher.Stop()
      namespaceWatcher = congress.StartNamespaceWatcher(kubeClient, extClient, config)
      namespaceCh = namespaceWatcher.ResultChan()
    }

    if doPolicyRestart {
      log.Println("Congress is restarting the policy watcher")
      doPolicyRestart = false
      policyWatcher.Stop()
      policyWatcher = congress.StartPolicyWatcher(extClient, config)
      policyCh = policyWatcher.ResultChan()
    }
  }
}
