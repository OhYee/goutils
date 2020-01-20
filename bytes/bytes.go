package bytes

//go:generate gcg ./template/data.json

import (
	"bytes"
	"io"
)

// Serializable could be serialization type
type Serializable interface {
	ToBytes() []byte
}

// NewBuffer initial a empty bytes.Buffer
func NewBuffer(b ...byte) *bytes.Buffer {
	return bytes.NewBuffer(b)
}

// ReadNBytes read n bytes from a io.Reader
func ReadNBytes(r io.Reader, n int) (b []byte, err error) {
	readLength := 0
	buf := bytes.NewBuffer([]byte{})
	for readLength < n {
		var l int
		temp := make([]byte, n-readLength)
		l, err = r.Read(temp)
		temp = temp[:l]
		if err != nil {
			return
		}
		buf.Write(temp)
		readLength += l
	}
	b = buf.Bytes()
	return
}

// FromBytesWithLength32 the data to the writer after the length of the data
// The length will store with int32
func FromBytesWithLength32(b []byte) (bb []byte) {
	bb = append(FromInt32(int32(len(b))), b...)
	return
}

// FromBytesWithLength64 the data to the writer after the length of the data
// The length will store with int64
func FromBytesWithLength64(b []byte) (bb []byte) {
	bb = append(FromInt64(int64(len(b))), b...)
	return
}

// ReadBytesWithLength32 read an int32 length first, and then read the data with length
func ReadBytesWithLength32(r io.Reader) (b []byte, err error) {
	l, err := ReadInt32(r)
	if err != nil {
		return
	}
	b, err = ReadNBytes(r, int(l))
	return
}

// ReadBytesWithLength64 read an int64 length first, and then read the data with length
// ONLY USE IN 64-BIT SYSTEM
func ReadBytesWithLength64(r io.Reader) (b []byte, err error) {
	l, err := ReadInt64(r)
	if err != nil {
		return
	}
	b, err = ReadNBytes(r, int(l))
	return
}
