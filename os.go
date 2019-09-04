package goutil

import (
	"strings"
)

// PathConcat concat Path
func PathConcat(path ...string) string {
	lst := make([]string, len(path))
	for idx, s := range path {
		s = strings.ReplaceAll(s, "\\\\", "/")
		if s[len(s)-1] == '/' {
			lst[idx] = s[0 : len(s)-1]
		} else {
			lst[idx] = s
		}
	}
	return strings.Join(lst, "/")
}
