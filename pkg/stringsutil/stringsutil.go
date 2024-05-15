package stringsutil

import (
	"strings"

	"github.com/joakim-ribier/go-utils/pkg/genericsutil"
)

type stringS string

// ##
// #### stringS type functions ####
// ##

// NewStringS builds {stringS} type from {s}.
func NewStringS(s string) stringS {
	return stringS(s)
}

// OrElse returns {s} if the value is not empty else returns {orIsEmpty}.
func (s stringS) OrElse(orIsEmpty string) string {
	return OrElse(s.S(), orIsEmpty)
}

// ReplaceAll replaces in {s} the {old} value by the {new} value.
func (s stringS) ReplaceAll(old, new string) stringS {
	return stringS(strings.ReplaceAll(string(s), old, new))
}

// S returns the string type of {s}.
func (s stringS) S() string {
	return string(s)
}

// When returns {isTrue} when the provided {cond} function returns true else returns {isFalse}.
func (s stringS) When(cond func(string) bool, isTrue, isFalse string) string {
	return genericsutil.When(string(s), cond, isTrue, isFalse)
}

// ##
// #### string type functions ####
// ##

// OrElse returns {s} if not empty else returns {orIsEmpty}.
func OrElse(s, orIsEmpty string) string {
	return genericsutil.OrElse(s, func() bool {
		return IsNotEmpty(s)
	}, orIsEmpty)
}

// IsEmpty returns true if {s} is empty else false.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty returns true if {s} is not empty else false.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}
