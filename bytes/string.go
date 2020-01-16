package bytes

import (
	"io"
)

// FromString transfer from string to []byte
func FromString(value string) (b []byte) {
	b = []byte(value)
	return
}

// FromStringWithLength32 transfer string to []byte with 32-bit length
func FromStringWithLength32(value string) (b []byte) {
	b = FromBytesWithLength32([]byte(value))
	return
}

// FromStringWithLength64 transfer string to []byte with 64-bit length
func FromStringWithLength64(value string) (b []byte) {
	b = FromBytesWithLength64([]byte(value))
	return
}

// ReadStringWithLength32 read string with 32-bit length from reader
func ReadStringWithLength32(r io.Reader) (value string, err error) {
	var b []byte
	b, err = ReadBytesWithLength32(r)
	if err == nil {
		value = string(b)
	}
	return
}

// ReadStringWithLength64 read string with 64-bit length from reader
func ReadStringWithLength64(r io.Reader) (value string, err error) {
	var b []byte
	b, err = ReadBytesWithLength64(r)
	if err == nil {
		value = string(b)
	}
	return
}
