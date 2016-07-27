package utils

// SearchStrings searchs a slice of strings for the given string and returns its index or -1
func SearchStrings(set []string, target string) int {
  for ndx, s := range set {
    if s == target {
      return ndx
    }
  }

  return -1
}