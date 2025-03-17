package dsa

import "errors"

type node struct {
	value any
	next  *node
}

type doublyLinkedList struct {
	head *node
	tail *node
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

func (ll *singlyLinkedList) append(value any) {
	tail := ll.tail()
	if tail != nil {
		tail.next = &node{value: value}
	} else {
		ll.head = &node{value: value}
	}
}

func (dll *doublyLinkedList) append(value any) {
	appendingNode := newNode(value)
	if dll.head == nil {
		dll.head = appendingNode
		dll.tail = appendingNode
	} else {
		dll.tail.next = appendingNode
		dll.tail = appendingNode
	}
}

func (dll *doublyLinkedList) removeFromStart() (*node, error) {
	if dll.head == nil {
		return nil, errors.New("list is empty")
	}
	removing := dll.head
	newHead := removing.next
	dll.head = newHead
	return removing, nil
}

func (ll *singlyLinkedList) search(value any) (*node, error) {
	curr := ll.head
	for {
		if curr == nil {
			return nil, errors.New("value not present in list")
		}
		if curr.value == value {
			return curr, nil
		}
		curr = curr.next
	}
}

func (ll *singlyLinkedList) insertAfter(c *node, value any) {
	inserting := newNode(value)
	next := c.next
	c.next = inserting
	inserting.next = next
}
