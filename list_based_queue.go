package dsa

import "errors"

type ListBasedQueue struct {
	storage *doublyLinkedList
}

func newListBasedQueue() *ListBasedQueue {
	return &ListBasedQueue{storage: &doublyLinkedList{}}
}

func (l ListBasedQueue) enqueue(a any) {
	l.storage.append(a)
}

func (l ListBasedQueue) dequeue() (any, error) {
	removedNode, err := l.storage.removeFromStart()
	if err != nil {
		return nil, errors.New("queue is empty")
	}
	return removedNode.value, nil
}

func (l ListBasedQueue) read() (any, error) {
	t := l.storage.tail
	if t == nil {
		return nil, errors.New("queue is empty")
	}
	return t.value, nil
}
