package dsa

import (
	"errors"
)

type node struct {
	value any
	next  *node
}
type singlyLinkedList struct {
	head *node
}
type linkedListStack struct {
	storage singlyLinkedList
}

func newLLStack() *linkedListStack {
	return &linkedListStack{storage: singlyLinkedList{}}
}

func (s *linkedListStack) push(el any) {
	if s.empty() {
		s.storage.head = &node{value: el}
	} else {
		s.storage.tail().next = &node{value: el}
	}
}

func (s *linkedListStack) pop() (any, error) {
	if s.empty() {
		return nil, errors.New("stack is empty")
	}
	var value any
	curr := s.storage.head
	for {
		if curr.next == nil {
			value = curr.value
			s.storage.head = nil
			//fmt.Println("reached break condition when stack is 1 element deep, popping off", value)
			break
		} else if curr.next.next == nil {
			value = curr.next.value
			curr.next = nil
			//fmt.Println("reached break condition, popping off", value)
			break
		}
		curr = curr.next
	}
	return value, nil
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

func (ll singlyLinkedList) tail() *node {
	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}
	return curr
}
