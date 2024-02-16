package iosutil

import (
	"os"
)

// Load loads data from a file {filename}.
func Load(filename string) ([]byte, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Write writes data in the file {filename}.
func Write(bytes []byte, filename string) error {
	return os.WriteFile(filename, bytes, 0644)
}
