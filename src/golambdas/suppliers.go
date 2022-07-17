package golambdas

func Supplier[T any](supply func() T) T {
	return supply()
}
