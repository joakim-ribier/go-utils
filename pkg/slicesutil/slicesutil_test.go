package slicesutil

import (
	"slices"
	"strconv"
	"testing"
	"time"
)

type UserTestMock struct {
	Age       int
	Name      string
	Hobbies   []string
	CreatedAt time.Time
}

// ##
// #### sliceS type functions ####
// ##

// TestAppend calls slicesutil.NewSliceS,
// checking for a valid return value.
func TestNewSliceS(t *testing.T) {
	_s := sliceS{"one", "three", "nine", "two"}
	r := NewSliceS(_s)
	if !Equal(_s, r) {
		t.Fatalf(`result: {%v} but expected: {%v}`, false, true)
	}
}

// TestAppend calls slicesutil.Append,
// checking for a valid return value.
func TestAppend(t *testing.T) {
	_s := sliceS{"one", "three", "nine", "two"}
	r := _s.Append(sliceS{"one", "four", "four"})

	if len(r) != 5 || r[len(r)-1] != "four" {
		t.Fatalf(`result: {%s} but expected: {%s}`, r, sliceS{"one", "three", "nine", "two", "four"})
	}
}

// TestClone calls slicesutil.Clone,
// checking for a valid return value.
func TestClone(t *testing.T) {
	_s := sliceS{"one", "three", "nine", "two"}
	r := _s.Clone()

	slices.Reverse(_s)

	if len(_s) != len(r) || _s[0] == r[0] {
		t.Fatalf(`result: {%s} but expected: {%s}`, r, _s)
	}
}

// TestDistinct calls slicesutil.Distinct,
// checking for a valid return value.
func TestDistinct(t *testing.T) {
	_s := sliceS{"one", "two", "three", "four", "one", "four"}

	r := _s.Distinct()

	if len(r) != 4 || !Must(r, sliceS{"one", "two", "three", "four"}) {
		t.Fatalf(`result: {%s} but expected: {%s}`, r, sliceS{"one", "two", "three", "four"})
	}
}

// TestExist calls slicesutil.Exist,
// checking for a valid return value.
func TestExist(t *testing.T) {
	_s := sliceS{"one", "three", "nine", "two"}
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
	_s := sliceS{"one", "three", "nine", "two"}
	r := _s.Filter(func(s string) bool { return s == "three" })

	if len(r) != 1 || r[0] != "three" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, sliceS{"three"})
	}
}

// TestFilterByNonEmpty calls slicesutil.FilterByNonEmpty,
// checking for a valid return value.
func TestFilterByNonEmpty(t *testing.T) {
	_s := sliceS{"", "", "", "two"}
	r := _s.FilterByNonEmpty()

	if len(r) != 1 || r[0] != "two" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "two")
	}
}

// TestFindLastOccurrenceIn calls slicesutil.FindLastOccurrenceIn,
// checking for a valid return value.
func TestFindLastOccurrenceIn(t *testing.T) {
	_s := sliceS{"one", "three", "nine", "two"}
	if r := _s.FindLastOccurrenceIn(sliceS{"nine", "one"}); r != "nine" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "nine")
	}
	if r := _s.FindLastOccurrenceIn(sliceS{"seven"}); r != "" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "")
	}
}

// TestFindNextEl calls slicesutil.FindNextEl,
// checking for a valid return value.
func TestFindNextEl(t *testing.T) {
	_s := sliceS{"one", "three", "nine", "two"}
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
	_s := sliceS{"z", "p", "a", "m", "a"}
	r := _s.Sort()

	if r[0] != "a" || r[1] != "a" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "a")
	}
	if r[len(_s)-1] != "z" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, "z")
	}
}

// TestToMap calls slicesutil.ToMap,
// checking for a valid return value.
func TestToMap(t *testing.T) {
	_s := sliceS{"one", "1", "two", "2"}
	r := _s.ToMap()

	if len(r) != 2 || r["one"] != "1" || r["two"] != "2" {
		t.Fatalf(`result: {%s} but expected: {%v}`, r, map[string]string{"one": "1", "two": "2"})
	}
}

// ##
// #### string functions ####
// ##

// TestContainAll calls slicesutil.ContainAll,
// checking for a valid return value.
func TestContainAll(t *testing.T) {
	_s1 := []string{"one", "three", "four", "two"}
	_s2 := []string{"three", "one", "two", "four", "five"}
	if !ContainAll(_s1, _s2) {
		t.Fatalf(`result: {false} but expected: {true}`)
	}
	if ContainAll(_s2, _s1) {
		t.Fatalf(`result: {true} but expected: {false}`)
	}
}

// TestEqual calls slicesutil.Equal,
// checking for a valid return value.
func TestEqual(t *testing.T) {
	_s1 := []string{"one", "three", "four", "two"}
	_s2 := []string{"three", "one", "two", "four"}
	if !Equal(_s1, _s2) {
		t.Fatalf(`result: {false} but expected: {true}`)
	}
	if Equal([]string{"one", "one"}, []string{"one", "two"}) {
		t.Fatalf(`result: {true} but expected: {false}`)
	}
	if Equal(_s1, append(_s2, "five")) {
		t.Fatalf(`result: {true} but expected: {false}`)
	}
}

// TestMust calls slicesutil.Must,
// checking for a valid return value.
func TestMust(t *testing.T) {
	_s1 := []string{"one", "two", "three", "four"}
	if !Must(_s1, slices.Clone(_s1)) {
		t.Fatalf(`result: {false} but expected: {true}`)
	}
	if Must(_s1, []string{"three", "one", "two", "four"}) {
		t.Fatalf(`result: {true} but expected: {false}`)
	}
	if Must(_s1, append(_s1, "five")) {
		t.Fatalf(`result: {true} but expected: {false}`)
	}
}

// ##
// #### generic functions ####
// ##

// TestAddOrReplaceT calls slicesutil.AddOrReplaceT,
// checking for a valid return value.
func TestAddOrReplaceT(t *testing.T) {
	_1 := []UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}}

	r := AddOrReplaceT(_1, UserTestMock{Name: "N-2", Age: 22}, func(utm UserTestMock) bool { return utm.Name == "N-2" })

	if len(r) != 2 || !ExistT(r, func(utm UserTestMock) bool { return utm.Name == "N-2" && utm.Age == 22 }) {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, []UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 22}})
	}
}

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

// TestFindT calls slicesutil.FindT,
// checking for a valid return value.
func TestFindT(t *testing.T) {
	_1 := []UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}, {Name: "N-3", Age: 2}}

	r := FindT(_1, func(arg0 UserTestMock) bool { return arg0.Age == 1 })
	if r.Name != "N-1" {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, []UserTestMock{{Name: "N-1", Age: 1}})
	}

	r = FindT(_1, func(arg0 UserTestMock) bool { return arg0.Age == 99 })
	if r != nil {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, "nil")
	}
}

// TestTransformT calls slicesutil.TransformT,
// checking for a valid return value.
func TestFlatTransformT(t *testing.T) {
	_1 := []UserTestMock{
		{Name: "N-1", Hobbies: []string{"playing chess", "coding"}},
		{Name: "N-2", Hobbies: []string{"running", "sailing", "cooking", "coding"}},
		{Name: "N-3"}}

	r := FlatTransformT[UserTestMock, string](_1, func(utm UserTestMock) ([]string, error) {
		return utm.Hobbies, nil
	})

	if len(r) != 6 {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, []string{"playing chess", "coding", "running", "sailing", "cooking", "coding"})
	}
}

// TestForAllT calls slicesutil.ForAllT,
// checking for a valid return value.
func TestForAllT(t *testing.T) {

	tcs := []struct {
		name      string
		value     []string
		predicate func(string) bool
		result    bool
	}{
		{
			name:      "with empty slices",
			value:     []string{},
			predicate: func(s string) bool { return false },
			result:    true,
		},
		{
			name:      "all values fulfill the predicate",
			value:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return len(s) == 1 },
			result:    true,
		},
		{
			name:      "not all values fulfill the predicate",
			value:     []string{"a", "b", "c"},
			predicate: func(s string) bool { return s == "a" },
			result:    false,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			_r := ForAllT(tc.value, tc.predicate)
			if _r != tc.result {
				t.Fatalf(`%s: {%v} but expected {%v}`, tc.name, _r, tc.result)
			}
		})
	}
}

// TestSortT calls slicesutil.SortT,
// checking for a valid return value.
func TestSortT(t *testing.T) {
	_1 := []UserTestMock{{Name: "N-3", Age: 6}, {Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}}

	r := SortT(_1, func(utm1, utm2 UserTestMock) (int, int) {
		return utm1.Age, utm2.Age
	})
	if r[0].Age != 1 || r[2].Age != 6 {
		t.Fatalf(`result: {%v} but expected: {%v}`, r,
			[]UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}, {Name: "N-3", Age: 6}})
	}
}

// TestSortTByTime calls slicesutil.SortTByTime,
// checking for a valid return value.
func TestSortTByTime(t *testing.T) {
	getTime := func(v string) time.Time {
		t, _ := time.Parse("2006-01-02", v)
		return t
	}

	_1 := []UserTestMock{
		{Name: "N-3", CreatedAt: getTime("2024-01-01")},
		{Name: "N-1", CreatedAt: getTime("2020-01-01")},
		{Name: "N-2", CreatedAt: getTime("2022-01-01")}}

	r := SortTByTime(_1, func(utm1, utm2 UserTestMock) (time.Time, time.Time) {
		return utm1.CreatedAt, utm2.CreatedAt
	})
	if r[0].Name != "N-1" || r[2].Name != "N-3" {
		t.Fatalf(`result: {%v} but expected: {%v}`, r,
			[]UserTestMock{
				{Name: "N-1", CreatedAt: getTime("2020-01-01")},
				{Name: "N-2", CreatedAt: getTime("2022-01-01")},
				{Name: "N-3", CreatedAt: getTime("2024-01-01")}})
	}

	r = SortTByTime(_1, func(utm1, utm2 UserTestMock) (time.Time, time.Time) {
		return utm2.CreatedAt, utm1.CreatedAt
	})
	if r[0].Name != "N-3" || r[2].Name != "N-1" {
		t.Fatalf(`result: {%v} but expected: {%v}`, r,
			[]UserTestMock{
				{Name: "N-3", CreatedAt: getTime("2024-01-01")},
				{Name: "N-2", CreatedAt: getTime("2022-01-01")},
				{Name: "N-1", CreatedAt: getTime("2020-01-01")}})
	}
}

// TestTransformT calls slicesutil.TransformT,
// checking for a valid return value.
func TestTransformT(t *testing.T) {
	_s := sliceS{"1", "2", "3", "4", "wrong-value"}

	r := TransformT[string, int](_s, func(s string) (*int, error) {
		v, err := strconv.Atoi(s)
		return &v, err
	})

	if r[0] != 1 || len(r) != 4 {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, []int{1, 2, 3, 4})
	}
}

// TestToString calls slicesutil.ToString,
// checking for a valid return value.
func TestToStringT(t *testing.T) {
	_1 := []UserTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}, {Name: "N-3", Age: 2}}

	r := ToStringT[UserTestMock](_1, func(u UserTestMock) *string { return &u.Name }, "#")

	expected := "N-1#N-2#N-3"
	if r != expected {
		t.Fatalf(`result: {%v} but expected: {%v}`, r, expected)
	}
}
