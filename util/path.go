package util

import (
	"path"
)

// PrefixPaths add prefixes from prefix to every path in paths
func PrefixPaths(prefix string, paths ...string) []string {
	ret := []string{}
	for _, s := range paths {
		ret = append(ret, path.Join(prefix, s))
	}
	return ret
}
