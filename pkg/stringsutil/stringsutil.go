package stringsutil

import "strings"

type stringS string

// ##
// #### stringS type functions ####
// ##

// NewStringS builds {stringS} type from {s}.
func NewStringS(s string) stringS {
	return stringS(s)
}

// OrElse returns {s} if the value is not empty else returns {or}.
func (s stringS) OrElse(or string) string {
	return orElse(s.S(), or)
}

// ReplaceAll replaces in {s} the {old} value by the {new} value.
func (s stringS) ReplaceAll(old, new string) stringS {
	return stringS(strings.ReplaceAll(string(s), old, new))
}

// S returns the string type of {s}.
func (s stringS) S() string {
	return string(s)
}

// When returns {is} when the provided {cond} function returns true else returns {or}.
func (s stringS) When(is, or string, cond func(string) bool) string {
	if cond(string(s)) {
		return is
	}
	return or
}

// ##
// #### string type functions ####
// ##

func orElse(s, or string) string {
	if IsEmpty(s) {
		return or
	}
	return s
}

// IsEmpty returns true if {s} is empty else false.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty returns true if {s} is not empty else false.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}
