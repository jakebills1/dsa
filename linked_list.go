package dsa

import "errors"

type node struct {
	value any
	next  *node
}
type singlyLinkedList struct {
	head *node
}

func (ll *singlyLinkedList) tail() *node {
	curr := ll.head
	for curr.next != nil {
		curr = curr.next
	}
	return curr
}

func newNode(value any) *node {
	return &node{value: value}
}

func (ll *singlyLinkedList) removeFromEnd() (any, error) {
	curr := ll.head
	if curr == nil {
		return nil, errors.New("no nodes in linked list yet")
	}
	for {
		if curr.next == nil {
			//fmt.Println("reached break condition when stack is 1 element deep, popping off", value)
			value := curr.value
			ll.head = nil
			return value, nil
		} else if curr.next.next == nil {
			//fmt.Println("reached break condition, popping off", value)
			value := curr.next.value
			curr.next = nil
			return value, nil
		}
		curr = curr.next
	}
}
