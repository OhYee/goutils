// FromFloat{{.bit}} transfer from float{{.bit}} to []byte
func FromFloat{{.bit}}(value float{{.bit}}) (b []byte) {
    b = make([]byte, {{.byte}})
    binary.BigEndian.PutUint{{.bit}}(b, math.Float{{.bit}}bits(value))
    return
}
