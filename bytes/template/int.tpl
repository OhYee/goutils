// FromUint{{.bit}} transfer from uint{{.bit}} to []byte
func FromUint{{.bit}}(value uint{{.bit}}) (b []byte) {
    b = make([]byte, {{.byte}})
    binary.BigEndian.PutUint{{.bit}}(b, uint{{.bit}}(value))
    return
}

// FromInt{{.bit}} transfer from int{{.bit}} to []byte
func FromInt{{.bit}}(value int{{.bit}}) (b []byte) {
    b = make([]byte, {{.byte}})
    binary.BigEndian.PutUint{{.bit}}(b, uint{{.bit}}(value))
    return
}