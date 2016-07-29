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

package utils

import (
  "fmt"
  "os"
  "strings"

  "k8s.io/kubernetes/pkg/labels"
)

const (
  // DefaultNamespaceExclude is the default list of namespaces to exclude in network isolation
  DefaultNamespaceExclude = "kube-system"
  // DefaultLabelSelector is the default label selector to use as a namespace watch filter
  DefaultLabelSelector = ""
  // DefaultIgnoreSelector is the default label selector to use to identify namespaces to be ignored
  DefaultIgnoreSelector = ""
  // DefaultIsolateNamespace is the default choice for locking down a namespace
  DefaultIsolateNamespace = false
)

// Config contains the configurable variables used by the congress policy maker and watcher
type Config struct {
  Excludes []string
  LabelSelector labels.Selector
  IgnoreSelector labels.Selector
  IsolateNamespace bool
}

// ConfigFromEnv returns the configuration based on the environment variables
func ConfigFromEnv() (*Config, error) {
  config := &Config{}

  label := os.Getenv("CONGRESS_SELECTOR")
  if label == "" {
    label = DefaultLabelSelector
  }

  selector, err := labels.Parse(label)
  if err != nil {
    return nil, fmt.Errorf("Failed to created label selector: %v", err)
  }

  config.LabelSelector = selector

  ignore := os.Getenv("CONGRESS_IGNORE_SELECTOR")
  if label == "" {
    label = DefaultLabelSelector
  }

  ignoreSelector, err := labels.Parse(ignore)
  if err != nil {
    return nil, fmt.Errorf("Failed to created label selector: %v", err)
  }

  config.IgnoreSelector = ignoreSelector

  excludesRaw := os.Getenv("CONGRESS_EXCLUDES")
  if excludesRaw == "" {
    excludesRaw = DefaultNamespaceExclude
  }

  config.Excludes = strings.Split(excludesRaw, ",")

  isolate := os.Getenv("CONGRESS_ISOLATE_NAMESPACE")
  if isolate == "true" {
    config.IsolateNamespace = true
  } else if isolate == "false" {
    config.IsolateNamespace = false
  } else {
    config.IsolateNamespace = DefaultIsolateNamespace
  }

  return config, nil
}

// IsExcluded checks to see if the given target namespace name should be excluded or not
func (c *Config) IsExcluded(target string) bool {
  if SearchStrings(c.Excludes, target) != -1 {
    // should be excluded from policies
    return true
  }

  // not excluded, write policies for it
  return false
}

// InIgnoreSelector checks to see if the given labels contain an ignored label.
func (c *Config) InIgnoreSelector(set map[string]string) bool {
  if set == nil {
    return false // not in the ignore selector
  }

  labelSet := labels.Set(set)
  return c.IgnoreSelector.Matches(labelSet)
}