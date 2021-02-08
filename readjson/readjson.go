package readjson

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

// ReaderToJSON parse a reader to a struct
func ReaderToJSON(r io.Reader, res interface{}) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, res); err != nil {
		return
	}
	return
}

// JSONToReader parse a struct to a reader
func JSONToReader(data interface{}) (r io.Reader, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	r = bytes.NewReader(b)
	return
}
