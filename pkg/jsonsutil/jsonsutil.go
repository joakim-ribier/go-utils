package jsonsutil

import (
	"encoding/json"
	"os"
)

func LoadF[T any](filename string) (T, error) {
	var data T

	file, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}

	return Unmarshal[T]([]byte(file))
}

func WriteT[T any](v T, filename string) error {
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

func Unmarshal[T any](bytes []byte) (T, error) {
	var data T

	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return data, err
	}

	return data, err
}
