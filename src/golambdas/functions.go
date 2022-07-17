package golambdas

func (s *BiStream[T, R]) FlatMap(apply func(element T) []R, list []T) []R {
	var result []R
	for _, element := range list {
		result = append(result, apply(element)...)
	}
	return result
}

func (s *BiStream[T, R]) Map(apply func(element T) R, list []T) []R {
	var result []R
	for _, element := range list {
		result = append(result, apply(element))
	}
	return result
}

func (s *BiStream[T, R]) Reduce(identity R, apply func(accumulator R, element T) R) R {
	for _, element := range s.Elements {
		identity = apply(identity, element)
	}
	return identity
}

func (s *Stream[T]) Reduce(identity T, apply func(accumulator T, element T) T) T {
	for _, element := range s.Elements {
		identity = apply(identity, element)
	}
	return identity
}
