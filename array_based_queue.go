package dsa

import (
	"errors"
)

type Queue interface {
	enqueue(any)
	dequeue() (any, error)
	read() (any, error)
}

type ArrayBasedQueue[T any] struct {
	storage []T
}

func (q *ArrayBasedQueue[T]) enqueue(el T) {
	q.storage = append(q.storage, el)
}

func (q *ArrayBasedQueue[T]) dequeue() (T, error) {
	if len(q.storage) == 0 {
		var zero T
		return zero, errors.New("Queue is empty")
	}
	el := q.storage[0]
	q.storage = q.storage[1:]
	return el, nil
}

func (q *ArrayBasedQueue[T]) read() (T, error) {
	if len(q.storage) == 0 {
		var zero T
		return zero, errors.New("Queue is empty")
	}
	el := q.storage[0]
	return el, nil
}

func newArrayQueue[T any]() *ArrayBasedQueue[T] {
	return &ArrayBasedQueue[T]{storage: []T{}}
}
