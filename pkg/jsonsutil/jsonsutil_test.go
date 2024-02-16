package jsonsutil

import (
	"testing"
)

type UserTestMock struct {
	Age  int
	Name string
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

// TestMarshal calls jsonsutil.Marshal,
// checking for a valid return value.
func TestMarshal(t *testing.T) {
	user := UserTestMock{Age: 18, Name: "Bob"}
	if bytes, err := Marshal(user); err != nil {
		t.Fatal(err)
	} else {
		if string(bytes) != `{"Age":18,"Name":"Bob"}` {
			t.Fatalf(`result: {%v} but expected: {%v}`, string(bytes), `{"Age":18,"Name":"Bob"}`)
		}
	}
}

// TestMarshalError calls jsonsutil.Marshal,
// checking for a valid return value.
func TestMarshalError(t *testing.T) {
	hello := func() string {
		return "Hello Word"
	}
	if bytes, err := Marshal(hello); err == nil {
		t.Fatalf(`result: {%v} but expected UnsupportedTypeError`, bytes)
	}
}
