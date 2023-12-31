package slicesutil

import (
	"slices"
	"sort"
	"strings"
)

type SliceS []string

// ##
// #### SliceS type functions ####
// ##

// Append appends the slice {to} into slice without duplicated values (case-insensitive).
func (s SliceS) Append(to []string) []string {
	return Append(s, to)
}

// Clone creates a new copy of slice.
func (s SliceS) Clone() []string {
	return Clone(s)
}

// Exist returns true if {v} value exists in slice (case-insensitive).
func (s SliceS) Exist(v string) bool {
	return Exist(s, v)
}

// Filter filters slice using the provided {is} function.
func (s SliceS) Filter(is func(string) bool) SliceS {
	return FilterT(s, is)
}

// FindLastOccurrenceIn finds last occurrence in slice of slice {to},
//
// returns empty if no occurrence found.
func (s SliceS) FindLastOccurrenceIn(to []string) string {
	return FindLastOccurrenceIn(s, to)
}

// FindNextEl finds the next element after the value {in} in the slice.
func (s SliceS) FindNextEl(in string) string {
	return FindNextEl(s, in)
}

// Sort sorts slice values.
func (s SliceS) Sort() []string {
	return Sort(s)
}

// ##
// #### string functions ####
// ##

// Append appends the slice {to} into the slice {from} without duplicated values (case-insensitive).
func Append(from, to []string) []string {
	return AppendT(from, to, func(f, t string) bool { return strings.EqualFold(f, t) })
}

// Exist returns {true} if {v} value exists in the slice {s} (case-insensitive).
func Exist(s []string, v string) bool {
	return ExistT(s, func(el string) bool { return strings.EqualFold(el, v) })
}

// FindLastOccurrenceIn finds last occurrence in the slice {from} of the slice {to},
//
// returns empty if no occurrence found.
func FindLastOccurrenceIn(from, to []string) string {
	s := Clone(from)
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
	sort.SliceStable(s, func(i, j int) bool {
		return strings.ToLower(s[i]) < strings.ToLower(s[j])
	})
	return s
}

// ##
// #### generic functions ####
// ##

// Append appends the slice {to} into the slice {from} without duplicated values using the provided {exist} function.
func AppendT[T any](from, to []T, exist func(T, T) bool) []T {
	s := Clone(from)
	for _, t := range to {
		if !ExistT(s, func(el T) bool { return exist(el, t) }) {
			s = append(s, t)
		}
	}
	return s
}

// Clone creates new copy of the slice {s}.
func Clone[T any](s []T) []T {
	return append([]T{}, s...)
}

// Exist returns bool using the provided equal function.
func ExistT[T any](s []T, equal func(T) bool) bool {
	return slices.IndexFunc(s, func(el T) bool { return equal(el) }) != -1
}

// Filter filters the slice in using the provided is function.
func FilterT[T any](in []T, is func(T) bool) []T {
	var out []T
	for _, v := range in {
		if is(v) {
			out = append(out, v)
		}
	}
	return out
}

// Transform transforms the slice {from} []F to the slice []T using the provider {transform} function.
func Transform[F any, T any](from []F, transorm func(F) (T, error)) []T {
	var to []T
	for _, el := range from {
		if v, err := transorm(el); err == nil {
			to = append(to, v)
		}
	}
	return to
}
