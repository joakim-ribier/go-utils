package mapsutil

import "testing"

type userTestMock struct {
	Age  int
	Name string
}

// TestSortT calls mapsutil.SortT,
// checking for a valid return value.
func TestSortT(t *testing.T) {
	_1 := map[userTestMock]string{
		{Name: "N-3", Age: 3}: "N-3",
		{Name: "N-1", Age: 1}: "N-1",
		{Name: "N-2", Age: 2}: "N-2",
	}

	r := SortT[userTestMock, string, int](_1, func(utm1, utm2 userTestMock) (int, int) {
		return utm1.Age, utm2.Age
	})

	if r[0] != "N-1" || r[1] != "N-2" || r[2] != "N-3" {
		t.Fatalf(`result: {%v} but expected: {%v}`, r,
			[]userTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}, {Name: "N-3", Age: 6}})
	}
}

// TestSort calls mapsutil.Sort,
// checking for a valid return value.
func TestSort(t *testing.T) {
	_1 := map[string]userTestMock{
		"N-3": {Name: "N-3", Age: 3},
		"N-1": {Name: "N-1", Age: 1},
		"N-2": {Name: "N-2", Age: 2},
	}

	r := Sort(_1)

	if r[0].Name != "N-1" || r[1].Name != "N-2" || r[2].Name != "N-3" {
		t.Fatalf(`result: {%v} but expected: {%v}`, r,
			[]userTestMock{{Name: "N-1", Age: 1}, {Name: "N-2", Age: 2}, {Name: "N-3", Age: 6}})
	}
}
