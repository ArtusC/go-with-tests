package generics

/*
Stacks should be fairly straightforward to understand from a requirements point of view.
They're a collection of items where you can Push items to the "top" and to get items
back again you Pop items from the top (LIFO - last in, first out).
*/

/*
In terms of constraints, any does mean "anything" and so does interface{}.
In fact, any was added in 1.18 and is just an alias for interface{}.
*/
type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zerto T
		return zerto, false
	}

	idx := len(s.values) - 1
	el := s.values[idx]

	s.values = s.values[:idx]

	return el, true

}
