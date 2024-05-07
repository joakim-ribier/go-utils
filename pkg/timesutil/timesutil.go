package timesutil

import "time"

// TWithExecutionTime struct contains execution time of {doSomething} function.
type TWithExecutionTime[T any] struct {
	T            *T
	TimeInMillis int64
}

// WithExecutionTime computes the execution time of the provided {doSomething} function.
func WithExecutionTime[T any](doSomething func() (*T, error)) (*TWithExecutionTime[T], error) {
	start := time.Now().UnixMilli()
	if resp, err := doSomething(); err != nil {
		return nil, err
	} else {
		end := time.Now().UnixMilli()
		return &TWithExecutionTime[T]{
			T: resp, TimeInMillis: end - start,
		}, nil
	}
}
