// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.8 (https://github.com/OhYee/gcg)

package fp

import (
	"reflect"
	"testing"
)

func TestMapString(t *testing.T) {
	type args struct {
		f     func(string, int) string
		input []string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []string
	}{
		{
			name:       "add a 1 at the end of the string",
			args:       args{f: func(s string, idx int) string { return s + "1" }, input: []string{"a", "bb", "ccc", "dddd"}},
			wantOutput: []string{"a1", "bb1", "ccc1", "dddd1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapString(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapInt(t *testing.T) {
	type args struct {
		f     func(int, int) int
		input []int
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n int, idx int) int { return n + 1 }, input: []int{1, 2, 3, 4, 5}},
			wantOutput: []int{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n int, idx int) int { return n - 1 }, input: []int{1, 2, 3, 4, 5}},
			wantOutput: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapInt(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapInt() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapInt8(t *testing.T) {
	type args struct {
		f     func(int8, int) int8
		input []int8
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int8
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n int8, idx int) int8 { return n + 1 }, input: []int8{1, 2, 3, 4, 5}},
			wantOutput: []int8{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n int8, idx int) int8 { return n - 1 }, input: []int8{1, 2, 3, 4, 5}},
			wantOutput: []int8{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapInt8(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapInt8() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapInt16(t *testing.T) {
	type args struct {
		f     func(int16, int) int16
		input []int16
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int16
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n int16, idx int) int16 { return n + 1 }, input: []int16{1, 2, 3, 4, 5}},
			wantOutput: []int16{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n int16, idx int) int16 { return n - 1 }, input: []int16{1, 2, 3, 4, 5}},
			wantOutput: []int16{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapInt16(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapInt16() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapInt32(t *testing.T) {
	type args struct {
		f     func(int32, int) int32
		input []int32
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int32
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n int32, idx int) int32 { return n + 1 }, input: []int32{1, 2, 3, 4, 5}},
			wantOutput: []int32{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n int32, idx int) int32 { return n - 1 }, input: []int32{1, 2, 3, 4, 5}},
			wantOutput: []int32{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapInt32(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapInt32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapInt64(t *testing.T) {
	type args struct {
		f     func(int64, int) int64
		input []int64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []int64
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n int64, idx int) int64 { return n + 1 }, input: []int64{1, 2, 3, 4, 5}},
			wantOutput: []int64{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n int64, idx int) int64 { return n - 1 }, input: []int64{1, 2, 3, 4, 5}},
			wantOutput: []int64{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapInt64(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapInt64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapUint8(t *testing.T) {
	type args struct {
		f     func(uint8, int) uint8
		input []uint8
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []uint8
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n uint8, idx int) uint8 { return n + 1 }, input: []uint8{1, 2, 3, 4, 5}},
			wantOutput: []uint8{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n uint8, idx int) uint8 { return n - 1 }, input: []uint8{1, 2, 3, 4, 5}},
			wantOutput: []uint8{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapUint8(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapUint8() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapUint16(t *testing.T) {
	type args struct {
		f     func(uint16, int) uint16
		input []uint16
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []uint16
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n uint16, idx int) uint16 { return n + 1 }, input: []uint16{1, 2, 3, 4, 5}},
			wantOutput: []uint16{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n uint16, idx int) uint16 { return n - 1 }, input: []uint16{1, 2, 3, 4, 5}},
			wantOutput: []uint16{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapUint16(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapUint16() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapUint32(t *testing.T) {
	type args struct {
		f     func(uint32, int) uint32
		input []uint32
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []uint32
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n uint32, idx int) uint32 { return n + 1 }, input: []uint32{1, 2, 3, 4, 5}},
			wantOutput: []uint32{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n uint32, idx int) uint32 { return n - 1 }, input: []uint32{1, 2, 3, 4, 5}},
			wantOutput: []uint32{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapUint32(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapUint32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapUint64(t *testing.T) {
	type args struct {
		f     func(uint64, int) uint64
		input []uint64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []uint64
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n uint64, idx int) uint64 { return n + 1 }, input: []uint64{1, 2, 3, 4, 5}},
			wantOutput: []uint64{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n uint64, idx int) uint64 { return n - 1 }, input: []uint64{1, 2, 3, 4, 5}},
			wantOutput: []uint64{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapUint64(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapUint64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapFloat32(t *testing.T) {
	type args struct {
		f     func(float32, int) float32
		input []float32
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []float32
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n float32, idx int) float32 { return n + 1 }, input: []float32{1, 2, 3, 4, 5}},
			wantOutput: []float32{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n float32, idx int) float32 { return n - 1 }, input: []float32{1, 2, 3, 4, 5}},
			wantOutput: []float32{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapFloat32(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapFloat32() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapFloat64(t *testing.T) {
	type args struct {
		f     func(float64, int) float64
		input []float64
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []float64
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n float64, idx int) float64 { return n + 1 }, input: []float64{1, 2, 3, 4, 5}},
			wantOutput: []float64{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n float64, idx int) float64 { return n - 1 }, input: []float64{1, 2, 3, 4, 5}},
			wantOutput: []float64{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapFloat64(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapFloat64() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMapByte(t *testing.T) {
	type args struct {
		f     func(byte, int) byte
		input []byte
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []byte
	}{
		{
			name:       "all value + 1",
			args:       args{f: func(n byte, idx int) byte { return n + 1 }, input: []byte{1, 2, 3, 4, 5}},
			wantOutput: []byte{2, 3, 4, 5, 6},
		},
		{
			name:       "all value - 1",
			args:       args{f: func(n byte, idx int) byte { return n - 1 }, input: []byte{1, 2, 3, 4, 5}},
			wantOutput: []byte{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := MapByte(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("MapByte() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		f     func(any, int) any
		input []any
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []any
	}{
		{
			name: "change all int value + 1 or output aaa",
			args: args{
				f: func(s any, idx int) any {
					if reflect.TypeOf(s).Kind() == reflect.Int {
						return s.(int) + 1
					}
					return "aaa"
				},
				input: []any{1, 2, "c", "d"},
			},
			wantOutput: []any{2, 3, "aaa", "aaa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := Map(tt.args.f, tt.args.input); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("Map() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
