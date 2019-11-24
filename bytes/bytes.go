package bytes

//go:generate gcg ./template/data.json

import ()

// FromString transfer from string to []byte
func FromString(value string) (b []byte) {
	b = []byte(value)
	return
}
