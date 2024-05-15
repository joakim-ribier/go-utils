package genericsutil

// OrElse returns {value} when the provided {cond} function returns true else returns {orIsFalse}.
func OrElse[T any](value T, cond func() bool, orIsFalse T) T {
	if cond() {
		return value
	} else {
		return orIsFalse
	}
}

// When returns {isTrue} when the provided {cond} function returns true else returns {isFalse} value.
func When[T any, R any](value T, cond func(T) bool, isTrue, isFalse R) R {
	if cond(value) {
		return isTrue
	}
	return isFalse
}
