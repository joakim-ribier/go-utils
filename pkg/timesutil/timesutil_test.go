package timesutil

import (
	"errors"
	"testing"
	"time"
)

// TestWithExecutionTime calls timesutil.WithExecutionTime,
// checking for a valid return value.
func TestWithExecutionTime(t *testing.T) {
	r, _ := WithExecutionTime[string](func() (*string, error) {
		time.Sleep(250 * time.Millisecond)
		var r string = "Hello World"
		return &r, nil
	})
	if *r.T != "Hello World" {
		t.Fatalf(`result: {%s} but expected: {%s}`, *r.T, "Hello World")
	}
	if uint64(r.TimeInMillis) < 250 || uint64(r.TimeInMillis) > 260 {
		t.Fatalf(`result: {%d} but expected: {%s}`, uint64(r.TimeInMillis), "between 500 and 505 ms")
	}
}

// TestWithExecutionTimeError calls timesutil.WithExecutionTime,
// checking for a valid return value.
func TestWithExecutionTimeError(t *testing.T) {
	_, err := WithExecutionTime[string](func() (*string, error) {
		time.Sleep(500 * time.Millisecond)
		return nil, errors.New("NPE error")
	})
	if err == nil {
		t.Fatalf(`result: {%s} but expected: {%s}`, "nil", err)
	}
}
