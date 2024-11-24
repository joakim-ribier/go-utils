package mapsutil

import (
	"cmp"
	"maps"
	"slices"

	"github.com/joakim-ribier/go-utils/pkg/slicesutil"
)

// Sort returns map {v} values sorted by map keys.
func Sort[R any](m map[string]R) []R {
	return sortT(m, slicesutil.Sort(slices.Collect(maps.Keys(m))))
}

// SortT returns map {v} values sorted by map keys using the provided {get} function.
func SortT[T comparable, R any, E cmp.Ordered](m map[T]R, get func(T, T) (E, E)) []R {
	return sortT(m, slicesutil.SortT[T, E](slices.Collect(maps.Keys(m)), get))
}

func sortT[T comparable, R any](m map[T]R, sortedKeys []T) []R {
	out := make([]R, 0, len(m))
	for _, b := range sortedKeys {
		out = append(out, m[b])
	}
	return out
}
