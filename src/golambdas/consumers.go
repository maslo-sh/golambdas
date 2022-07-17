package golambdas

func (s *Stream[T]) ForEach(accept func(element T)) {
	for _, element := range s.Elements {
		accept(element)
	}
}
