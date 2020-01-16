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
	b = FromString(value)
	b = append(FromInt32(int32(len(b))), b...)
	return
}

// FromStringWithLength64 transfer string to []byte with 64-bit length
func FromStringWithLength64(value string) (b []byte) {
	b = FromString(value)
	b = append(FromInt64(int64(len(b))), b...)
	return
}

// ReadStringWithLength32 read string with 32-bit length from reader
func ReadStringWithLength32(r io.Reader) (value string, err error) {
	l, err := ReadInt32(r)
	if err != nil {
		return
	}

	b, err := ReadNBytes(r, int(l))
	if err != nil {
		return
	}

	value = string(b)
	return
}

// ReadStringWithLength64 read string with 64-bit length from reader
func ReadStringWithLength64(r io.Reader) (value string, err error) {
	l, err := ReadInt64(r)
	if err != nil {
		return
	}

	b, err := ReadNBytes(r, int(l))
	if err != nil {
		return
	}

	value = string(b)
	return
}
