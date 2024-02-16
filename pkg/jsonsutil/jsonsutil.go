package jsonsutil

import (
	"encoding/json"
)

// Marshal returns the JSON encoding of {v}.
func Marshal[T any](t T) ([]byte, error) {
	bytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Unmarshal parses the JSON-encoded {bytes} and returns {T}.
func Unmarshal[T any](bytes []byte) (T, error) {
	var data T
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
