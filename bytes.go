package goutil

import (
	"bytes"
	// "fmt"
	"io"
)

// ReadNBytes read n bytes from a io.Reader
func ReadNBytes(r io.Reader, n int) (b []byte, err error) {
	readLength := 0
	buf := bytes.NewBuffer([]byte{})
	for readLength < n {
		var l int
		temp := make([]byte, n-readLength)
		l, err = r.Read(temp)
		temp=temp[:l]
		if err != nil {
			return
		}
		buf.Write(temp)
		readLength += l
	}
	b = buf.Bytes()
	return
}
