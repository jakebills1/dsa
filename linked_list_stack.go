package dsa

import (
	"errors"
)

type linkedListStack struct {
	storage singlyLinkedList
}

func newLLStack() *linkedListStack {
	return &linkedListStack{storage: singlyLinkedList{}}
}

func (s *linkedListStack) push(el any) {
	if s.empty() {
		s.storage.head = newNode(el)
	} else {
		s.storage.tail().next = newNode(el)
	}
}

func (s *linkedListStack) pop() (any, error) {
	value, err := s.storage.removeFromEnd()
	if err != nil {
		return nil, errors.New("stack is empty")
	}
	return value, err
}

func (s *linkedListStack) read() (any, error) {
	if s.empty() {
		return nil, errors.New("stack is empty")
	}
	return s.storage.tail().value, nil
}

func (s *linkedListStack) empty() bool {
	if s.storage.head == nil {
		return true
	}
	return false
}
