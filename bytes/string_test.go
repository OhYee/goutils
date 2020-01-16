package bytes

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestFromStringWithLength32(t *testing.T) {
	tests := []struct {
		name  string
		value string
		wantB []byte
	}{
		{
			name:  "abc",
			value: "abc",
			wantB: []byte{0x00, 0x00, 0x00, 0x03, 97, 98, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := FromStringWithLength32(tt.value); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("FromStringWithLength32() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestFromStringWithLength64(t *testing.T) {
	tests := []struct {
		name  string
		value string
		wantB []byte
	}{
		{
			name:  "abc",
			value: "abc",
			wantB: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 97, 98, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := FromStringWithLength64(tt.value); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("FromStringWithLength64() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestReadStringWithLength32(t *testing.T) {
	tests := []struct {
		name      string
		r         io.Reader
		wantValue string
		wantErr   bool
	}{
		{
			name:      "abc",
			r:         bytes.NewBuffer([]byte{0x00, 0x00, 0x00, 0x03, 97, 98, 99}),
			wantValue: "abc",
			wantErr:   false,
		},
		{
			name:      "length error",
			r:         bytes.NewBuffer([]byte{0x00}),
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "string error",
			r:         bytes.NewBuffer([]byte{0x00, 0x00, 0x00, 0x03, 97, 98}),
			wantValue: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := ReadStringWithLength32(tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadStringWithLength32() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ReadStringWithLength32() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestReadStringWithLength64(t *testing.T) {
	tests := []struct {
		name      string
		r         io.Reader
		wantValue string
		wantErr   bool
	}{
		{
			name:      "abc",
			r:         bytes.NewBuffer([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 97, 98, 99}),
			wantValue: "abc",
			wantErr:   false,
		},
		{
			name:      "length error",
			r:         bytes.NewBuffer([]byte{0x00}),
			wantValue: "",
			wantErr:   true,
		},
		{
			name:      "string error",
			r:         bytes.NewBuffer([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 97, 98}),
			wantValue: "",
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, err := ReadStringWithLength64(tt.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadStringWithLength64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValue != tt.wantValue {
				t.Errorf("ReadStringWithLength64() = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
