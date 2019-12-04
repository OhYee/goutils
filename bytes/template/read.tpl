// ReadInt{{.bit}} read an int{{.bit}} value
func ReadInt{{.bit}}(r io.Reader) (v int{{.bit}}, err error) {
	b, err := ReadNBytes(r, {{.byte}})
	if err != nil {
		return
	}
	v = ToInt{{.bit}}(b)
	return
}

// ReadUint{{.bit}} read an uint{{.bit}} value
func ReadUint{{.bit}}(r io.Reader) (v uint{{.bit}}, err error) {
	b, err := ReadNBytes(r, {{.byte}})
	if err != nil {
		return
	}
	v = ToUint{{.bit}}(b)
	return
}
