// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.8 (https://github.com/OhYee/gcg)

package bytes

import (
	"bytes"
	"github.com/OhYee/goutils/compare"
	"testing"
)

func TestReadUint8(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    uint8
		wantErr bool
	}{
		{
			name:    "test uint8 25",
			b:       []byte{0x19},
			want:    25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadUint8(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint8 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadUint16(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    uint16
		wantErr bool
	}{
		{
			name:    "test uint16 25*25",
			b:       []byte{0x02, 0x71},
			want:    25 * 25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadUint16(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint16 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadUint32(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    uint32
		wantErr bool
	}{
		{
			name:    "test uint32 25*25*25*25",
			b:       []byte{0x00, 0x05, 0xf5, 0xe1},
			want:    25 * 25 * 25 * 25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadUint32(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint32 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadUint64(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    uint64
		wantErr bool
	}{
		{
			name:    "test uint64 25*25*25*25*25*25*25*25",
			b:       []byte{0x00, 0x00, 0x00, 0x23, 0x86, 0xf2, 0x6f, 0xc1},
			want:    25 * 25 * 25 * 25 * 25 * 25 * 25 * 25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadUint64(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadUint64 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadInt8(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    int8
		wantErr bool
	}{
		{
			name:    "test int8 25",
			b:       []byte{0x19},
			want:    25,
			wantErr: false,
		},

		{
			name:    "test int8 -25",
			b:       []byte{0xE7},
			want:    -25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInt8(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt8 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadInt16(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    int16
		wantErr bool
	}{
		{
			name:    "test int16 25*25",
			b:       []byte{0x02, 0x71},
			want:    25 * 25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInt16(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt16 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadInt32(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    int32
		wantErr bool
	}{
		{
			name:    "test int32 25*25*25*25",
			b:       []byte{0x00, 0x05, 0xf5, 0xe1},
			want:    25 * 25 * 25 * 25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInt32(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt32 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadInt64(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    int64
		wantErr bool
	}{
		{
			name:    "test int64 25*25*25*25*25*25*25*25",
			b:       []byte{0x00, 0x00, 0x00, 0x23, 0x86, 0xf2, 0x6f, 0xc1},
			want:    25 * 25 * 25 * 25 * 25 * 25 * 25 * 25,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadInt64(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadInt64 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadFloat32(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    float32
		wantErr bool
	}{
		{
			name:    "test float32 1.0",
			b:       []byte{0x3f, 0x80, 0x00, 0x00},
			want:    1.0,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFloat32(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFloat32 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReadFloat64(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    float64
		wantErr bool
	}{
		{
			name:    "test float64 1.0",
			b:       []byte{0x3f, 0xf0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
			want:    1.0,
			wantErr: false,
		},

		{
			name:    "test empty bytes",
			b:       []byte{},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFloat64(bytes.NewBuffer(tt.b))
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFloat64 want error %v, but got %v\n", tt.wantErr, err)
			}
			if !tt.wantErr && !compare.Equal(got, tt.want) {
				t.Errorf("ReadFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}
