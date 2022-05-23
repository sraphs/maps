package maps

import (
	"bytes"
	"encoding/gob"
)

// https://gist.github.com/soroushjp/0ec92102641ddfc3ad5515ca76405f4d
func init() {
	gob.Register(map[string]interface{}{})
}

// Clone performs a deep copy of the given map m.
func Clone(m map[string]interface{}) (map[string]interface{}, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		return nil, err
	}
	var copy map[string]interface{}
	err = dec.Decode(&copy)
	if err != nil {
		return nil, err
	}
	return copy, nil
}
