package stack

type node[T any] struct {
	next  *node[T]
	value T
}

type Stack[T any] struct {
	head *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		head: nil,
	}
}

func (s *Stack[T]) Push(item T) {
	newHead := &node[T]{
		next:  s.head,
		value: item,
	}
	s.head = newHead
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.head != nil {
		item := s.head.value
		s.head = s.head.next
		return item, true
	}

	var t T
	return t, false
}

func (s *Stack[T]) Head() (T, bool) {
	if s.head != nil {
		return s.head.value, true
	}

	var t T
	return t, false
}
