package stringsutil

import (
	"testing"
)

// TestAppend calls stringsutil.OrElse,
// checking for a valid return value.
func TestStringOrElse(t *testing.T) {
	tcs := []struct {
		name   string
		value  string
		orElse string
		result string
	}{
		{
			name:   "with value",
			value:  "Hello",
			orElse: "no-value",
			result: "Hello",
		},
		{
			name:   "with empty value",
			value:  "",
			orElse: "empty-value",
			result: "empty-value",
		},
		{
			name:   "with space value",
			value:  "   ",
			orElse: "empty-value",
			result: "empty-value",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := NewStringS(tc.value).OrElse(tc.orElse)
			if _r != tc.result {
				t.Fatalf(`result: {%s} but expected: {%s}`, _r, tc.result)
			}
		})
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
func TestStringWhen(t *testing.T) {
	tcs := []struct {
		name   string
		value  string
		is     string
		or     string
		cond   func(string) bool
		result string
	}{
		{
			name:   "with value",
			value:  "1",
			is:     "true",
			or:     "false",
			cond:   func(s string) bool { return s == "1" },
			result: "true",
		},
		{
			name:   "with empty value",
			value:  " ",
			is:     "true",
			or:     "false",
			cond:   func(s string) bool { return s == "1" },
			result: "false",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := NewStringS(tc.value).When(tc.is, tc.or, tc.cond)
			if _r != tc.result {
				t.Fatalf(`result: {%s} but expected: {%s}`, _r, tc.result)
			}
		})
	}
}
