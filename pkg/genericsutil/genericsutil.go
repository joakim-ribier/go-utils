package genericsutil

// OrElse returns {value} when the provided {predicate} function returns true else returns {orIsFalse}.
func OrElse[T any](value T, predicate func(T) bool, orIsFalse T) T {
	if predicate(value) {
		return value
	}
	return orIsFalse
}

// When returns {isTrue} when the provided {predicate} function returns true else returns {isFalse} value.
func When[T any, R any](value T, predicate func(T) bool, isTrue, isFalse R) R {
	if predicate(value) {
		return isTrue
	}
	return isFalse
}

// ForAll returns {true} if each value respects the provided {predicate] function.
func ForAll[T any](predicate func(T) bool, in ...T) bool {
	is := true
	for _, v := range in {
		is := predicate(v)
		if !is {
			return false
		}
	}
	return is
}
