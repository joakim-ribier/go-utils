package slicesutil

import (
	"slices"
	"strconv"
	"testing"
)

type UserTestMock struct {
	Age  int
	Name string
}

// ##
// #### SliceS type functions ####
// ##

// TestAppend calls slicesutil.Append,
// checking for a valid return value.
func TestAppend(t *testing.T) {
	_s := SliceS{"one", "three", "nine", "two"}
	r := _s.Append(SliceS{"one", "four"})

	if len(r) != 5 || r[len(r)-1] != "four" {
		t.Fatalf(`result: {%s} but expected: {%s}`, r, SliceS{"one", "three", "nine", "two", "four"})
	}
}

// TestClone calls slicesutil.Clone,
// checking for a valid return value.
func TestClone(t *testing.T) {
	_s := SliceS{"one", "three", "nine", "two"}
	r := _s.Clone()

	slices.Reverse(_s)

	if len(_s) != len(r) || _s[0] == r[0] {
		t.Fatalf(`result: {%s} but expected: {%s}`, r, _s)
	}
}

// TestExist calls slicesutil.Exist,
// checking for a valid return value.
func TestExist(t *testing.T) {
	_s := SliceS{"one", "three", "nine", "two"}
	if !_s.Exist("nine") {
		t.Fatalf(`result: {%v} but expected: {%v}`, false, true)
	}
	if _s.Exist("seven") {
		t.Fatalf(`result: {%v} but expected: {%v}`, true, false)
	}
}

// TestFilter calls slicesutil.Filter,
// checking for a valid return value.
func TestFilter(t *testing.T) {
	_s := SliceS{"one", "three", "nine", "two"}
	r := _s.Filter(func(s string) bool { return s == "three" })

	if len(r) != 1 || r[0] != "three" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, SliceS{"three"})
	}
}

// TestFindLastOccurrenceIn calls slicesutil.FindLastOccurrenceIn,
// checking for a valid return value.
func TestFindLastOccurrenceIn(t *testing.T) {
	_s := SliceS{"one", "three", "nine", "two"}
	if r := _s.FindLastOccurrenceIn(SliceS{"nine", "one"}); r != "nine" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "nine")
	}
	if r := _s.FindLastOccurrenceIn(SliceS{"seven"}); r != "" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "")
	}
}

// TestFindNextEl calls slicesutil.FindNextEl,
// checking for a valid return value.
func TestFindNextEl(t *testing.T) {
	_s := SliceS{"one", "three", "nine", "two"}
	if r := _s.FindNextEl("three"); r != "nine" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "nine")
	}
	if r := _s.FindNextEl("two"); r != "" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "")
	}
}

// TestFindNextEl calls slicesutil.FindNextEl,
// checking for a valid return value.
func TestSort(t *testing.T) {
	_s := SliceS{"one", "three", "nine", "two"}
	r := _s.Sort()

	if r[0] != "nine" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "nine")
	}
	if r[len(_s)-1] != "two" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "two")
	}
}

// ##
// #### generic functions ####
// ##

// TestAppendT calls slicesutil.AppendT,
// checking for a valid return value.
func TestAppendT(t *testing.T) {
	_1 := []UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}}
	_2 := []UserTestMock{{Name: "N-2", Age: 22}, {Name: "N-3", Age: 3}}

	r := AppendT(_1, _2, func(arg0, arg1 UserTestMock) bool { return arg0.Name == arg1.Name })

	if len(r) != 3 {
		t.Fatalf(`result: {%v} but expected: {%v}`, r,
			[]UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}, {Name: "N-3", Age: 3}})
	}
}

// TestExistT calls slicesutil.ExistT,
// checking for a valid return value.
func TestExistT(t *testing.T) {
	_1 := []UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}}

	if !ExistT(_1, func(arg0 UserTestMock) bool { return arg0.Name == "N-2" }) {
		t.Fatalf(`result: {%v} but expected: {%v}`, true, false)
	}
	if ExistT(_1, func(arg0 UserTestMock) bool { return arg0.Name == "N-3" }) {
		t.Fatalf(`result: {%v} but expected: {%v}`, false, true)
	}
}

// TestFilterT calls slicesutil.FilterT,
// checking for a valid return value.
func TestFilterT(t *testing.T) {
	_1 := []UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}, {Name: "N-3", Age: 2}}

	r := FilterT(_1, func(arg0 UserTestMock) bool { return arg0.Age == 1 })

	if r[0].Name != "N-1" || len(r) != 1 {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, []UserTestMock{{Name: "N-1", Age: 1}})
	}
}

// TestTransform calls slicesutil.TestTransform,
// checking for a valid return value.
func TestTransform(t *testing.T) {
	_s := SliceS{"1", "2", "3", "4", "wrong-value"}

	r := Transform[string, int](_s, func(s string) (int, error) {
		return strconv.Atoi(s)
	})

	if r[0] != 1 || len(r) != 4 {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, []int{1, 2, 3, 4})
	}
}
