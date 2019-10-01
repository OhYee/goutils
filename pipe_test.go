package goutils

import (
	"bytes"
	"io"
	"testing"
)

func TestPipe_Read(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		p    Pipe
	}{
		{
			name: "compare with bytes.buffer",
			r:    bytes.NewBuffer([]byte{0x01, 0x02, 0x03, 0x04, 0x05}),
			p:    NewPipe(16, []byte{0x01, 0x02, 0x03, 0x04, 0x05}),
		},
		{
			name: "compare with bytes.buffer - part of data",
			r:    bytes.NewBuffer([]byte{0x01}),
			p:    NewPipe(16, []byte{0x01}),
		},
		{
			name: "EOF",
			r:    bytes.NewBuffer([]byte{}),
			p:    NewPipe(16, []byte{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data1 := make([]byte, 2, 4)
			data2 := make([]byte, 2, 4)
			data1[0] = 0x00
			data2[0] = 0x00
			tt.p.Close()
			n1, err1 := tt.r.Read(data1)
			n2, err2 := tt.p.Read(data2)
			if !Equal(n1, n2) || !Equal(err1, err2) || !Equal(data1, data2) {
				t.Errorf("%v %v %v != %v %v %v", n1, err1, data1, n2, err2, data2)
			} else {
				t.Logf("%v %v %v == %v %v %v", n1, err1, data1, n2, err2, data2)
			}
		})
	}
}

func TestPipe_Write(t *testing.T) {
	t.Run("write error", func(t *testing.T) {
		p := NewPipe(16, []byte{0x01})
		p.Close()
		n, err := p.Write([]byte{0x02})
		b := make([]byte, 16)
		n2, err2 := p.Read(b)
		b = b[:n2]
		if n != 0 || !Equal(err, errwrite) || n2 != 1 || err2 != nil || !Equal(b, []byte{0x01}) {
			t.Errorf("n=%d err=%v n2=%d err2=%v bytes=%v", n, err, n2, err2, b)
		}
	})
}
