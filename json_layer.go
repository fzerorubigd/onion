package onion

import (
	"encoding/json"
	"io"
	"os"
)

// NewJSONLayer creates a layer based on json stream
func NewJSONLayer(r io.Reader) (Layer, error) {
	dec := json.NewDecoder(r)
	var data map[string]interface{}
	err := dec.Decode(&data)
	if err != nil {
		return nil, err
	}

	return NewMapLayer(data), nil
}

// NewJSONFileLayer read a json file and create a layer based on it
func NewJSONFileLayer(f string) (Layer, error) {
	fl, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer fl.Close()

	return NewJSONLayer(fl)
}
