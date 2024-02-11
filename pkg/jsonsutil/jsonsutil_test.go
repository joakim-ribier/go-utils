package jsonsutil

import (
	"os"
	"testing"
)

type UserTestMock struct {
	Age  int
	Name string
}

// TestUnmarshalWithBadJson calls jsonsutil.Unmarshal,
// checking for a valid return value.
func TestLoadFWithBadFilename(t *testing.T) {
	r, err := LoadF[UserTestMock]("file-does-not-exist.json")
	if err == nil {
		t.Fatalf(`result: {%v} but expected error`, r)
	}
}

// TestUnmarshalWithBadJson calls jsonsutil.Unmarshal,
// checking for a valid return value.
func TestUnmarshalWithBadJson(t *testing.T) {
	r, err := Unmarshal[UserTestMock]([]byte("{broken json data...}"))
	if err == nil {
		t.Fatalf(`result: {%v} but expected error`, r)
	}
}

// TestUnmarshal calls jsonsutil.Unmarshal,
// checking for a valid return value.
func TestUnmarshal(t *testing.T) {
	user := UserTestMock{Age: 18, Name: "Bob"}
	if r, err := Unmarshal[UserTestMock]([]byte("{\"Age\": 18, \"Name\": \"Bob\"}")); err == nil {
		if r != user {
			t.Fatalf(`result: {%v} but expected: {%v}`, r, user)
		}
	} else {
		t.Fatal(err)
	}
}

// TestWriteTAndLoadF calls jsonsutil.WriteT and jsonsutil.LoadF,
// checking for a valid return value.
func TestWriteTAndLoadF(t *testing.T) {
	user := UserTestMock{Age: 18, Name: "Bob"}
	if err := WriteT(user, "users.json"); err != nil {
		t.Fatal(err)
	}
	if r, err := LoadF[UserTestMock]("users.json"); err == nil {
		if r != user {
			t.Fatalf(`result: {%v} but expected: {%v}`, r, user)
		}
		os.Remove("users.json")
	} else {
		t.Fatal(err)
	}
}
