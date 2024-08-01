package stringsutil

import (
	"strconv"
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

// Bool converts {in} string to boolean value
func Bool(in string) bool {
	return genericsutil.When(in, func(arg string) bool {
		return arg == "true" || arg == "1"
	}, true, false)
}

// Int converts {in} string to int value
func Int(in string, or int) int {
	if v, err := strconv.Atoi(in); err == nil {
		return v
	}
	return or
}

// IsEmpty returns true if {s} is empty else false.
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// IsNotEmpty returns true if {s} is not empty else false.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// OrElse returns {s} if not empty else returns {orIsEmpty}.
func OrElse(s, orIsEmpty string) string {
	return genericsutil.OrElse(s, func() bool {
		return IsNotEmpty(s)
	}, orIsEmpty)
}

// Split slices {value} into all substrings separated by {sep} not enclosed by {enclose} character
// and returns a slice of the substrings between those separators.
//
// Ex: Split(`Tom Hello Sam "How are you?" Tom Fine!, " ", `"`)
// => ["Tom", "Hello", "Sam", "How are you?", "Tom", "Fine!"]
func Split(value, sep, enclose string) []string {
	quote := false
	tab := []string{}
	el := ""
	for i, r := range strings.TrimSpace(value) {
		add := false
		concat := false
		lastRune := ""
		if i > 0 {
			lastRune = string(value[i-1])
		}
		character := string(r)
		if (character == enclose && lastRune != "\\") || quote {
			if !quote {
				quote = true
			} else {
				add = (character == enclose && lastRune != "\\") && quote
				concat = true
			}
		} else {
			if character == sep {
				add = character == sep
			} else {
				concat = true
			}
		}
		if add {
			if len(el) > 0 && el != enclose && el != sep {
				tab = append(tab, el)
			}
			el = ""
			quote = false
		}
		if concat && (el != enclose || lastRune != "\\") && character != "\\" {
			el = el + character
		}
	}
	if len(el) > 0 && el != enclose {
		tab = append(tab, el)
	}
	return tab
}
