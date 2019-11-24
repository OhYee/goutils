// Code generated by GCG. DO NOT EDIT.
// Go Code Generator 0.0.7 (https://github.com/OhYee/gcg)

package bytes

import (
	"encoding/binary"
	"math"
)

// FromFloat32 transfer from float32 to []byte
func FromFloat32(value float32) (b []byte) {
	b = make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(value))
	return
}

// FromFloat64 transfer from float64 to []byte
func FromFloat64(value float64) (b []byte) {
	b = make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(value))
	return
}
