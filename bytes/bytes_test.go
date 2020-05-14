package bytes

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/OhYee/goutils/compare"
	"github.com/OhYee/goutils/pipeline"
)

func TestReadNBytes(t *testing.T) {
	testCases := []struct {
		name    string
		input   []byte
		length  int
		want    []byte
		wantErr error
	}{
		{
			name:    "simple",
			input:   []byte{0x01, 0x02, 0x03},
			length:  3,
			want:    []byte{0x01, 0x02, 0x03},
			wantErr: nil,
		},
		{
			name:    "read part of data",
			input:   []byte{0x01, 0x02, 0x03},
			length:  2,
			want:    []byte{0x01, 0x02},
			wantErr: nil,
		},
		{
			name:    "read more than data",
			input:   []byte{0x01, 0x02, 0x03},
			length:  4,
			want:    []byte{},
			wantErr: fmt.Errorf("EOF"),
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			b, err := ReadNBytes(bytes.NewBuffer(testCase.input), testCase.length)
			if !bytes.Equal(b, testCase.want) || 0 != strings.Compare(fmt.Sprintf("%v", err), fmt.Sprintf("%v", testCase.wantErr)) {
				t.Errorf("Except %v %v, but got %v %v", testCase.want, testCase.wantErr, b, err)
			}
		})
	}
	t.Run("dynamic add data", func(t *testing.T) {
		pipe := pipeline.New(16, []byte{0x01, 0x02})
		answer := []byte{0x01, 0x02, 0x03}
		var b []byte
		var err error
		c := make(chan bool, 8)
		go func() {
			b, err = ReadNBytes(pipe, 3)
			c <- true
		}()
		go func() {
			<-time.After(time.Second * 1)
			pipe.Write([]byte{0x03})
			pipe.Close()
		}()

		<-c
		if !compare.Equal(b, answer) || !compare.Equal(err, nil) {
			t.Errorf("Except %v %v, but got %v %v", answer, nil, b, err)
		}
	})
}

func TestFromBytesWithLength32(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want []byte
	}{
		{
			name: "0x1234",
			b:    []byte{0x12, 0x34},
			want: []byte{0, 0, 0, 2, 0x12, 0x34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB := FromBytesWithLength32(tt.b)
			if !compare.Equal(gotB, tt.want) {
				t.Errorf("Want %v, got %v", tt.want, gotB)
			}
		})
	}
}

func TestFromBytesWithLength64(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want []byte
	}{
		{
			name: "0x1234",
			b:    []byte{0x12, 0x34},
			want: []byte{0, 0, 0, 0, 0, 0, 0, 2, 0x12, 0x34},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB := FromBytesWithLength64(tt.b)
			if !compare.Equal(gotB, tt.want) {
				t.Errorf("Want %v, got %v", tt.want, gotB)
			}
		})
	}
}

func TestReadWithLength32(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    []byte
		wantErr bool
	}{
		{
			name:    "0x1234",
			b:       []byte{0, 0, 0, 2, 0x12, 0x34},
			want:    []byte{0x12, 0x34},
			wantErr: false,
		},
		{
			name:    "do not have enough length",
			b:       []byte{},
			want:    []byte{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.b)
			got, err := ReadBytesWithLength32(buf)
			if (err == nil) == tt.wantErr {
				t.Errorf("err is %v, but want %v\n", err, tt.wantErr)
			} else if !compare.Equal(got, tt.want) {
				t.Errorf("Want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestReadWithLength64(t *testing.T) {
	tests := []struct {
		name    string
		b       []byte
		want    []byte
		wantErr bool
	}{
		{
			name:    "0x1234",
			b:       []byte{0, 0, 0, 0, 0, 0, 0, 2, 0x12, 0x34},
			want:    []byte{0x12, 0x34},
			wantErr: false,
		},
		{
			name:    "do not have enough length",
			b:       []byte{},
			want:    []byte{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.NewBuffer(tt.b)
			got, err := ReadBytesWithLength64(buf)
			if (err == nil) == tt.wantErr {
				t.Errorf("err is %v, but want %v\n", err, tt.wantErr)
			} else if !compare.Equal(got, tt.want) {
				t.Errorf("Want %v, got %v", tt.want, got)
			}
		})
	}
}

func TestNewBuffer(t *testing.T) {
	t.Run("buffer", func(t *testing.T) {
		buf := NewBuffer()
		buf.Write([]byte{0x01})
		if !compare.Equal(buf.Bytes(), []byte{0x01}) {
			t.Errorf("NewBuffer() = %v, want 0x01", buf.Bytes())
		}
	})
}
