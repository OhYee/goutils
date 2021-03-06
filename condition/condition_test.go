// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.8 (https://github.com/OhYee/gcg)

package condition

import (
	"reflect"
	"testing"
)

func TestIf(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          any
		f          any
		wantOutput any
	}{
		{
			name:       "any true",
			condition:  true,
			t:          "t",
			f:          "f",
			wantOutput: "t",
		},
		{
			name:       "any false",
			condition:  false,
			t:          "t",
			f:          "f",
			wantOutput: "f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := If(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("If() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfRune(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          rune
		f          rune
		wantOutput rune
	}{
		{
			name:       "rune true",
			condition:  true,
			t:          'a',
			f:          'b',
			wantOutput: 'a',
		},
		{
			name:       "rune false",
			condition:  false,
			t:          'a',
			f:          'b',
			wantOutput: 'b',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfRune(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfRune() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfString(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          string
		f          string
		wantOutput string
	}{
		{
			name:       "string true",
			condition:  true,
			t:          "a",
			f:          "b",
			wantOutput: "a",
		},
		{
			name:       "string false",
			condition:  false,
			t:          "a",
			f:          "b",
			wantOutput: "b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfString(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfByte(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          byte
		f          byte
		wantOutput byte
	}{
		{
			name:       "byte true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "byte false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfByte(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfByte() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfBool(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          bool
		f          bool
		wantOutput bool
	}{
		{
			name:       "bool true",
			condition:  true,
			t:          true,
			f:          false,
			wantOutput: true,
		},
		{
			name:       "bool false",
			condition:  false,
			t:          true,
			f:          false,
			wantOutput: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfBool(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfBool() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfFunction(t *testing.T) {
	var res bool
	tests := []struct {
		name       string
		condition  bool
		t          function
		f          function
		wantOutput function
	}{
		{
			name:       "function true",
			condition:  true,
			t:          func() { res = true },
			f:          func() { res = false },
			wantOutput: func() { res = true },
		},
		{
			name:       "function false",
			condition:  false,
			t:          func() { res = true },
			f:          func() { res = false },
			wantOutput: func() { res = false },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res = false
			IfFunction(tt.condition, tt.t, tt.f)()
			if res != tt.condition {
				t.Errorf("IfFunction() = %v, want %v", res, tt.condition)
			}

		})
	}
}
func TestIfInt(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          int
		f          int
		wantOutput int
	}{
		{
			name:       "int true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "int false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfInt(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfInt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfInt8(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          int8
		f          int8
		wantOutput int8
	}{
		{
			name:       "int8 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "int8 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfInt8(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfInt8() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfInt16(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          int16
		f          int16
		wantOutput int16
	}{
		{
			name:       "int16 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "int16 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfInt16(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfInt16() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfInt32(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          int32
		f          int32
		wantOutput int32
	}{
		{
			name:       "int32 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "int32 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfInt32(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfInt32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfInt64(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          int64
		f          int64
		wantOutput int64
	}{
		{
			name:       "int64 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "int64 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfInt64(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfInt64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfUint8(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          uint8
		f          uint8
		wantOutput uint8
	}{
		{
			name:       "uint8 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "uint8 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfUint8(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfUint8() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfUint16(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          uint16
		f          uint16
		wantOutput uint16
	}{
		{
			name:       "uint16 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "uint16 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfUint16(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfUint16() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfUint32(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          uint32
		f          uint32
		wantOutput uint32
	}{
		{
			name:       "uint32 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "uint32 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfUint32(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfUint32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfUint64(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          uint64
		f          uint64
		wantOutput uint64
	}{
		{
			name:       "uint64 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "uint64 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfUint64(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfUint64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfFloat32(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          float32
		f          float32
		wantOutput float32
	}{
		{
			name:       "float32 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "float32 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfFloat32(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfFloat32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestIfFloat64(t *testing.T) {
	tests := []struct {
		name       string
		condition  bool
		t          float64
		f          float64
		wantOutput float64
	}{
		{
			name:       "float64 true",
			condition:  true,
			t:          1,
			f:          2,
			wantOutput: 1,
		},
		{
			name:       "float64 false",
			condition:  false,
			t:          1,
			f:          2,
			wantOutput: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := IfFloat64(tt.condition, tt.t, tt.f); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("IfFloat64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
