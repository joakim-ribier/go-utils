package slicesutil

import (
	"cmp"
	"slices"
	"strings"
	"time"
)

type sliceS []string

// ##
// #### SliceS type functions ####
// ##

// NewSliceS builds {sliceS} type from {in}.
func NewSliceS(in []string) sliceS {
	return append(sliceS{}, in...)
}

// Append appends the slice {to} into slice without duplicated values (case-insensitive).
func (s sliceS) Append(to []string) sliceS {
	return Append(s, to)
}

// Clone creates a new copy of slice.
func (s sliceS) Clone() sliceS {
	return slices.Clone(s)
}

// Distinct removes duplicated values in the slice.
func (s sliceS) Distinct() sliceS {
	return Distinct(s)
}

// Exist returns true if {v} value exists in slice (case-insensitive).
func (s sliceS) Exist(v string) bool {
	return Exist(s, v)
}

// Filter filters slice using the provided {is} function.
func (s sliceS) Filter(is func(string) bool) sliceS {
	return FilterT(s, is)
}

// FilterByNonEmpty removes all empty values in the slice {in}.
func (s sliceS) FilterByNonEmpty() sliceS {
	return FilterByNonEmpty(s)
}

// FindLastOccurrenceIn finds last occurrence in slice of slice {to},
//
// returns empty if no occurrence found.
func (s sliceS) FindLastOccurrenceIn(to []string) string {
	return FindLastOccurrenceIn(s, to)
}

// FindNextEl finds the next element after the value {in} in the slice.
func (s sliceS) FindNextEl(in string) string {
	return FindNextEl(s, in)
}

// Sort sorts slice values.
func (s sliceS) Sort() []string {
	return Sort(s)
}

// ToMap transforms the slice {in} to a map key/value.
func (s sliceS) ToMap() map[string]string {
	return ToMap(s)
}

// ##
// #### string functions ####
// ##

// Append appends the slice {to} into the slice {from} without duplicated values (case-insensitive).
func Append(from, to []string) []string {
	return AppendT(from, to, func(f, t string) bool { return strings.EqualFold(f, t) })
}

// ContainAll returns {true} if every value of the slice {s1} is in the slice {s2} (case-insensitive).
func ContainAll(s1, s2 []string) bool {
	for _, v := range s1 {
		if !Exist(s2, v) {
			return false
		}
	}
	return true
}

// Distinct removes duplicated values in the slice {in}.
func Distinct(in []string) []string {
	return Append([]string{}, in)
}

// Equal returns {true} if the two slices contain exactly the same values (case-insensitive).
func Equal(s1, s2 []string) bool {
	return len(s1) == len(s2) && ContainAll(s1, s2) && ContainAll(s2, s1)
}

// Exist returns {true} if {v} value exists in the slice {s} (case-insensitive).
func Exist(s []string, v string) bool {
	return ExistT(s, func(el string) bool { return strings.EqualFold(el, v) })
}

// FilterByNonEmpty removes all empty values in the slice {in}.
func FilterByNonEmpty(in []string) []string {
	return FilterT(in, func(s string) bool { return strings.TrimSpace(s) != "" })
}

// Must returns {true} if the two slices contain exactly in the same order the same values (case-insensitive).
func Must(s1, s2 []string) bool {
	return slices.EqualFunc(s1, s2, strings.EqualFold)
}

// ToMap transforms the slice {in} to a map key/value.
func ToMap(in []string) map[string]string {
	out := make(map[string]string)
	for i := range in {
		if len(in) > i*2+1 {
			out[in[i*2]] = in[i*2+1]
		}
	}
	return out
}

// FindLastOccurrenceIn finds last occurrence in the slice {from} of the slice {to},
//
// returns empty if no occurrence found.
func FindLastOccurrenceIn(from, to []string) string {
	s := slices.Clone(from)
	slices.Reverse(s)

	for _, v := range s {
		if slices.Contains(to, v) {
			return v
		}
	}
	return ""
}

// FindNextEl finds the next element after the value {v} in the slice {s}.
func FindNextEl(s []string, v string) string {
	for i, el := range s {
		if el == v && len(s) > i+1 {
			return s[i+1]
		}
	}
	return ""
}

// Sort sorts the slice {s} values.
func Sort(s []string) []string {
	return SortT(s, func(a, b string) (string, string) {
		return strings.ToLower(a), strings.ToLower(b)
	})
}

// ##
// #### generic functions ####
// ##

// AddOrReplaceT adds or replaces {t} in the slice {in} using the provided {is} function.
func AddOrReplaceT[T any](in []T, t T, is func(T) bool) []T {
	var out []T
	for _, v := range in {
		if !is(v) {
			out = append(out, v)
		}
	}
	out = append(out, t)
	return out
}

// Append appends the slice {to} into the slice {from} without duplicated values using the provided {exist} function.
func AppendT[T any](from, to []T, exist func(T, T) bool) []T {
	s := slices.Clone(from)
	for _, t := range to {
		if !ExistT(s, func(el T) bool { return exist(el, t) }) {
			s = append(s, t)
		}
	}
	return s
}

// ExistT returns bool using the provided equal function.
func ExistT[T any](s []T, equal func(T) bool) bool {
	return slices.IndexFunc(s, func(el T) bool { return equal(el) }) != -1
}

// FilterT filters the slice in using the provided is function.
func FilterT[T any](in []T, is func(T) bool) []T {
	var out []T
	for _, v := range in {
		if is(v) {
			out = append(out, v)
		}
	}
	return out
}

// FindT finds the first occurrence in the slice in using the provided {is} function.
func FindT[T any](in []T, is func(T) bool) *T {
	if out := FilterT[T](in, is); len(out) > 0 {
		return &out[0]
	} else {
		return nil
	}
}

// FlatTransformT transforms and flattens the slice {from} []F to the slice []T using the provided {transform} function.
func FlatTransformT[F any, T any](from []F, transform func(F) ([]T, error)) []T {
	var to []T
	for _, el := range from {
		if v, err := transform(el); err == nil && v != nil {
			to = append(to, v...)
		}
	}
	return to
}

// ForAllT returns true if all the elements fulfil the predicate.
func ForAllT[T any](in []T, predicate func(T) bool) bool {
	is := true
	for _, v := range in {
		is := predicate(v)
		if !is {
			return false
		}
	}
	return is
}

// SortT clones and sorts the slice {s} using the provided {get} function.
func SortT[T any, E cmp.Ordered](s []T, get func(T, T) (E, E)) []T {
	var out = slices.Clone(s)
	slices.SortStableFunc(out, func(a, b T) int {
		f1, f2 := get(a, b)
		switch {
		case f1 < f2:
			return -1
		case f1 > f2:
			return +1
		}
		return 0
	})
	return out
}

// SortTByTime clones and sorts the slice {s} by time.Time type using the provided {get} function.
func SortTByTime[T any](s []T, get func(T, T) (time.Time, time.Time)) []T {
	var out = slices.Clone(s)
	slices.SortStableFunc(out, func(a, b T) int {
		t1, t2 := get(a, b)
		return t1.Compare(t2)
	})
	return out
}

// TransformT transforms the slice {from} []F to the slice []T using the provided {transform} function.
func TransformT[F any, T any](from []F, transform func(F) (*T, error)) []T {
	var to []T
	for _, el := range from {
		if v, err := transform(el); err == nil && v != nil {
			to = append(to, *v)
		}
	}
	return to
}

// ToString concatenates the slice {in} to a single string using the provided {transform} function.
func ToStringT[T any](in []T, transform func(T) *string, separator string) string {
	return strings.Join(TransformT[T, string](in, func(t T) (*string, error) {
		return transform(t), nil
	}), separator)
}
