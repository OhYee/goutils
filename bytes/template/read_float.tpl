// ReadFloat{{.bit}} read a float{{.bit}} value
func ReadFloat{{.bit}}(r io.Reader) (v float{{.bit}}, err error) {
	b, err := ReadNBytes(r, {{.byte}})
	if err != nil {
		return
	}
	v = ToFloat{{.bit}}(b)
	return
}
