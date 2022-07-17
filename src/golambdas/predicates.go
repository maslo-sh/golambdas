package golambdas

func (s *Stream[T]) Filter(predicate func(element T) bool) []T {
	var result []T
	for _, element := range s.Elements {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return result
}

func (s *Stream[T]) AllMatch(predicate func(element T) bool) bool {
	var result []T
	for _, element := range s.Elements {
		if predicate(element) {
			result = append(result, element)
		}
	}
	return len(result) == len(s.Elements)
}
