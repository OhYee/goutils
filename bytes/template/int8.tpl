// FromUint8 transfer from uint8 to []byte
func FromUint8(value uint8) (b []byte) {
    b = []byte{uint8(value)}
    return
}

// FromInt8 transfer from int8 to []byte
func FromInt8(value int8) (b []byte) {
    b = []byte{uint8(value)}
    return
}

// ToUint8 transfer from []byte to uint8
func ToUint8(b []byte) (value uint8) {
    value = uint8(b[0])
    return
}

// ToInt8 transfer from []byte to int8
func ToInt8(b []byte) (value int8) {
    value = int8(b[0])
    return
}
