package stringsutil

import (
	"testing"
)

// TestAppend calls stringsutil.OrElse,
// checking for a valid return value.
func TestOrElse(t *testing.T) {
	_r := NewStringS("Hello").OrElse("GoodBye")
	if _r != "Hello" {
		t.Fatalf(`result: {%s} but expected: {%s}`, _r, "Hello")
	}

	_r = NewStringS("").OrElse("GoodBye")
	if _r != "GoodBye" {
		t.Fatalf(`result: {%s} but expected: {%s}`, _r, "GoodBye")
	}

	_r = NewStringS("  ").OrElse("GoodBye")
	if _r != "GoodBye" {
		t.Fatalf(`result: {%s} but expected: {%s}`, _r, "GoodBye")
	}
}

// TestAppend calls stringsutil.ReplaceAll,
// checking for a valid return value.
func TestReplaceAll(t *testing.T) {
	_r := NewStringS("one two three four").
		ReplaceAll("one", "1").
		ReplaceAll("three", "3").S()

	if _r != "1 two 3 four" {
		t.Fatalf(`result: {%s} but expected: {%s}`, _r, "1 two 3 four")
	}
}

// TestAppend calls stringsutil.When,
// checking for a valid return value.
func TestWhen(t *testing.T) {
	_r := NewStringS("1").When("true", "false", func(s string) bool {
		return s == "1"
	})
	if _r != "true" {
		t.Fatalf(`result: {%s} but expected: {%s}`, _r, "true")
	}

	_r = NewStringS("Boby").When("Hello World", "Hello Boby", IsEmpty)
	if _r != "Hello Boby" {
		t.Fatalf(`result: {%s} but expected: {%s}`, _r, "Hello Boby")
	}

	_r = NewStringS("").When("Hello Boby", "Hello World", IsNotEmpty)
	if _r != "Hello World" {
		t.Fatalf(`result: {%s} but expected: {%s}`, _r, "Hello World")
	}
}
