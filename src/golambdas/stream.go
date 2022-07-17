package golambdas

type Stream[T any] struct {
	Elements []T
}

type BiStream[T, R any] struct {
	Elements []T
}

func NewStream[T any](elements []T) *Stream[T] {
	return &Stream[T]{elements}
}
