package dsa

import "errors"

type stackable interface {
	push(any)
	pop() (any, error)
	read() (any, error)
	empty() bool
}

type arrayStack[T any] struct {
	storage []T
}

func newArrayStack[T any]() *arrayStack[T] {
	return &arrayStack[T]{storage: []T{}}
}

func (s *arrayStack[T]) push(el T) {
	s.storage = append(s.storage, el)
}

func (s *arrayStack[T]) pop() (T, error) {
	if s.lastIdx() < 0 {
		var zero T
		return zero, errors.New("stack is empty")
	}
	el := s.storage[s.lastIdx()]
	s.storage = s.storage[:s.lastIdx()]
	return el, nil
}

func (s *arrayStack[T]) read() (any, error) {
	if len(s.storage) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.storage[s.lastIdx()], nil
}

func (s *arrayStack[T]) empty() bool {
	if len(s.storage) == 0 {
		return true
	}
	return false
}

func (s *arrayStack[T]) lastIdx() int {
	return len(s.storage) - 1
}
