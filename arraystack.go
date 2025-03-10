package dsa

import "errors"

type stackable interface {
	push(any)
	pop() (any, error)
	read() (any, error)
	empty() bool
}

type arrayStack struct {
	storage []any
}

func newArrayStack() *arrayStack {
	var s []any
	return &arrayStack{storage: s}
}

func (s *arrayStack) push(el any) {
	s.storage = append(s.storage, el)
}

func (s *arrayStack) pop() (any, error) {
	if s.lastIdx() < 0 {
		return nil, errors.New("stack is empty")
	}
	el := s.storage[s.lastIdx()]
	s.storage = s.storage[:s.lastIdx()]
	return el, nil
}

func (s *arrayStack) read() (any, error) {
	if len(s.storage) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.storage[s.lastIdx()], nil
}

func (s *arrayStack) empty() bool {
	if len(s.storage) == 0 {
		return true
	}
	return false
}

func (s *arrayStack) lastIdx() int {
	return len(s.storage) - 1
}
