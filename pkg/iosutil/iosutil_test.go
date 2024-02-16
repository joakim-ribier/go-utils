package iosutil

import (
	"os"
	"testing"
)

// TestLoadWithBadFilename calls iosutil.Load,
// checking for a valid return value.
func TestLoadWithBadFilename(t *testing.T) {
	r, err := Load("file-does-not-exist.json")
	if err == nil {
		t.Fatalf(`result: {%v} but expected error`, r)
	}
}

// TestWrite calls iosutil.Write,
// checking for a valid return value.
func TestWrite(t *testing.T) {
	expected := "Hello World"
	if err := Write([]byte(expected), "helloworld.json"); err != nil {
		t.Fatal(err)
	}
	if r, err := Load("helloworld.json"); err != nil {
		t.Fatal(err)
	} else {
		if string(r) != expected {
			t.Fatalf(`result: {%v} but expected: {%v}`, r, expected)
		}
	}
	os.Remove("helloworld.json")
}
