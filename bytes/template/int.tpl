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

// ToUint{{.bit}} transfer from []byte to uint{{.bit}} 
func ToUint{{.bit}}(b []byte) (value uint{{.bit}}) {
    value = binary.BigEndian.Uint{{.bit}}(b)
    return
}

// ToInt{{.bit}} transfer from []byte to int{{.bit}}
func ToInt{{.bit}}(b []byte) (value int{{.bit}}) {
    value = int{{.bit}}(binary.BigEndian.Uint{{.bit}}(b))
    return
}
