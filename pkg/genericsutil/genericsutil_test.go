package genericsutil

import (
	"testing"
)

// ##
// #### generic type functions ####
// ##

// TestAppend calls genericsutil.OrElse,
// checking for a valid return value.
func TestOrElse(t *testing.T) {
	tcs := []struct {
		name   string
		value  string
		or     string
		cond   bool
		result string
	}{
		{
			name:   "return value if cond = true",
			value:  "1",
			or:     "2",
			cond:   true,
			result: "1",
		},
		{
			name:   "return or if cond = false",
			value:  "1",
			or:     "2",
			cond:   false,
			result: "2",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := OrElse(tc.value, func() bool { return tc.cond }, tc.or)
			if _r != tc.result {
				t.Fatalf(`result: {%v} but expected: {%v}`, _r, tc.result)
			}
		})
	}
}

// TestAppend calls genericsutil.When,
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
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := When(tc.value, tc.cond, tc.is, tc.or)
			if _r != tc.result {
				t.Fatalf(`result: {%s} but expected: {%s}`, _r, tc.result)
			}
		})
	}
}
