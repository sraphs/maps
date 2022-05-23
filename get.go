package maps

import (
	"errors"
	"strings"
)

var (
	// ErrNotFound is key not found.
	ErrNotFound = errors.New("key not found")
)

// Get read Value in given map[string]interface{}
// by the given path, will return false if not found.
func Get(values map[string]interface{}, path string) Value {
	var (
		next = values
		keys = strings.Split(path, ".")
		last = len(keys) - 1
	)
	for idx, key := range keys {
		value, ok := next[key]
		if !ok {
			return &errValue{err: ErrNotFound}
		}
		if idx == last {
			if value == nil {
				return &nilValue{}
			}
			av := &atomicValue{}
			av.Store(value)
			return av
		}
		switch vm := value.(type) {
		case map[string]interface{}:
			next = vm
		default:
			return &errValue{err: ErrNotFound}
		}
	}
	return &errValue{err: ErrNotFound}
}
