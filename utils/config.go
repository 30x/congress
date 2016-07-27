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
)

// Config contains the configurable variables used by the congress policy maker and watcher
type Config struct {
  Excludes []string
  LabelSelector labels.Selector
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

  excludesRaw := os.Getenv("CONGRESS_EXCLUDES")
  if excludesRaw == "" {
    excludesRaw = DefaultNamespaceExclude
  }

  config.Excludes = strings.Split(excludesRaw, ",")

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