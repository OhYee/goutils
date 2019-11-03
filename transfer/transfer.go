package transfer

import (
	"reflect"
)

type any = interface{}

// ToInterfaceSlice transfer any slice to []interface{}.
// useful in fmt.Sprintf("...", ToInterfaceSlice([]string{...}))
func ToInterfaceSlice(v any) (s []any) {
	kind := reflect.TypeOf(v).Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		slice := reflect.ValueOf(v)
		s = make([]any, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			s[i] = slice.Index(i).Interface()
		}
	} else {
		s = make([]any, 1)
		s[0] = v
	}
	return
}
