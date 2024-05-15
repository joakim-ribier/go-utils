package stringsutil

import (
	"testing"
)

// ##
// #### string type functions ####
// ##

// TestAppend calls stringsutil.IsEmpty,
// checking for a valid return value.
func TestIsEmpty(t *testing.T) {
	tcs := []struct {
		name   string
		value  string
		result bool
	}{
		{
			name:   "not empty",
			value:  "not empty",
			result: false,
		},
		{
			name:   "with empty value",
			value:  "",
			result: true,
		},
		{
			name:   "with space value",
			value:  "   ",
			result: true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := IsEmpty(tc.value)
			if _r != tc.result {
				t.Fatalf(`result: {%v} but expected: {%v}`, _r, tc.result)
			}
		})
	}
}

// TestAppend calls stringsutil.IsNotEmpty,
// checking for a valid return value.
func TestIsNotEmpty(t *testing.T) {
	tcs := []struct {
		name   string
		value  string
		result bool
	}{
		{
			name:   "not empty",
			value:  "not empty",
			result: true,
		},
		{
			name:   "with empty value",
			value:  "",
			result: false,
		},
		{
			name:   "with space value",
			value:  "   ",
			result: false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := IsNotEmpty(tc.value)
			if _r != tc.result {
				t.Fatalf(`result: {%v} but expected: {%v}`, _r, tc.result)
			}
		})
	}
}

// ##
// #### NewStringS type functions ####
// ##

// TestAppend calls stringsutil.OrElse,
// checking for a valid return value.
func TestOrElse(t *testing.T) {
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
func TestWhen(t *testing.T) {
	tcs := []struct {
		name   string
		value  string
		is     string
		or     string
		cond   func(string) bool
		result string
	}{
		{
			name:   "convert int to bool",
			value:  "1",
			is:     "true",
			or:     "false",
			cond:   func(s string) bool { return s == "1" },
			result: "true",
		},
		{
			name:   "convert int to bool with none value",
			value:  " ",
			is:     "true",
			or:     "false",
			cond:   func(s string) bool { return s == "1" },
			result: "false",
		},
		{
			name:   "with empty value function",
			value:  " ",
			is:     "is empty",
			or:     "not empty",
			cond:   IsEmpty,
			result: "is empty",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := NewStringS(tc.value).When(tc.cond, tc.is, tc.or)
			if _r != tc.result {
				t.Fatalf(`result: {%s} but expected: {%s}`, _r, tc.result)
			}
		})
	}
}
