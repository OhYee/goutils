package compare

import (
	"fmt"
	"strings"

	"github.com/OhYee/goutils"
)

// Equal two data using fmt
func Equal(a goutils.Any, b goutils.Any) bool {
	return strings.Compare(fmt.Sprintf("%v", a), fmt.Sprintf("%v", b)) == 0
}
