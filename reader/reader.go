package reader

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadJSON(r io.Reader, res interface{}) (err error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, res); err != nil {
		return
	}
	return
}
