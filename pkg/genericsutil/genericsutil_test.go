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
		name      string
		value     int
		or        int
		predicate func(int) bool
		result    int
	}{
		{
			name:      "should be 5 if value > 0",
			value:     5,
			or:        5,
			predicate: func(v int) bool { return v > 0 },
			result:    5,
		},
		{
			name:      "should be 5 if value < 0",
			value:     -5,
			or:        5,
			predicate: func(v int) bool { return v > 0 },
			result:    -5 * -1,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := OrElse(tc.value, tc.predicate, tc.or)
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
		name      string
		value     string
		is        string
		or        string
		predicate func(string) bool
		result    string
	}{
		{
			name:      `"1" should be true`,
			value:     "1",
			is:        "true",
			or:        "false",
			predicate: func(s string) bool { return s == "1" },
			result:    "true",
		},
		{
			name:      `" " should be false`,
			value:     " ",
			is:        "true",
			or:        "false",
			predicate: func(s string) bool { return s == "1" },
			result:    "false",
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := When(tc.value, tc.predicate, tc.is, tc.or)
			if _r != tc.result {
				t.Fatalf(`result: {%s} but expected: {%s}`, _r, tc.result)
			}
		})
	}
}

// TestForAll calls genericsutil.ForAll,
// checking for a valid return value.
func TestForAll(t *testing.T) {
	tcs := []struct {
		name      string
		value     []any
		predicate func(any) bool
		result    bool
	}{
		{
			name:      "should be true if each value equals true",
			value:     []any{true, true, true},
			predicate: func(a any) bool { return a.(bool) },
			result:    true,
		},
		{
			name:      "should be false if at least one value equals false",
			value:     []any{true, false, true},
			predicate: func(a any) bool { return a.(bool) },
			result:    false,
		},
		{
			name:      "should be true if each value is empty",
			value:     []any{"", ""},
			predicate: func(a any) bool { return a == "" },
			result:    true,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := ForAll[any](tc.predicate, tc.value...)
			if _r != tc.result {
				t.Fatalf(`result: {%v} but expected: {%v}`, _r, tc.result)
			}
		})
	}
}
