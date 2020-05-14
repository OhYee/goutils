package transfer

import (
	"testing"

	"github.com/OhYee/goutils/compare"
)

func TestToInterfaceSlice(t *testing.T) {

	tests := []struct {
		name string
		v    any
		want []any
	}{
		{
			name: "transfer slice to []interface{}",
			v:    []string{"a", "b", "c"},
			want: []interface{}{"a", "b", "c"},
		},
		{
			name: "transfer array to []interface{}",
			v:    [3]string{"a", "b", "c"},
			want: []interface{}{"a", "b", "c"},
		},
		{
			name: "transfer other type to []interface{}",
			v:    "a",
			want: []interface{}{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := ToInterfaceSlice(tt.v); !compare.Equal(gotS, tt.want) {
				t.Errorf("ToInterfaceSlice() = %v, want %v", gotS, tt.want)
			}
		})
	}
}
