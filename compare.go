package goutil

import (
	"fmt"
	"strings"
)

// Equal two data using fmt
func Equal(a Any, b Any) bool {
	return strings.Compare(fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)) == 0
}
