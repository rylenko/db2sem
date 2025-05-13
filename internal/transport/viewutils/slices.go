package viewutils

import "slices"

func ContainsString(slice []string, item string) bool {
	return slices.Contains(slice, item)
}
