// FromFloat{{.bit}} transfer from float{{.bit}} to []byte
func FromFloat{{.bit}}(value float{{.bit}}) (b []byte) {
    b = make([]byte, {{.byte}})
    binary.BigEndian.PutUint{{.bit}}(b, math.Float{{.bit}}bits(value))
    return
}

// ToFloat{{.bit}} transfer from []byte to float{{.bit}}
func ToFloat{{.bit}}(b []byte) (value float{{.bit}}) {
    value = math.Float{{.bit}}frombits(ToUint{{.bit}}(b))
    return
}
