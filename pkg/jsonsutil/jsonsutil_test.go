package jsonsutil

import (
	"os"
	"testing"
)

type UserTestMock struct {
	Age  int
	Name string
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
