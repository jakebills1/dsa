package dsa

import (
	"testing"
)

func Test_ArrayBasedQueue(t *testing.T) {
	subject := newArrayQueue[int]()
	subject.enqueue(1)
	subject.enqueue(2)
	subject.enqueue(3)
	testCases := []int{1, 2, 3}
	for _, n := range testCases {
		el, err := subject.dequeue()
		if el != n {
			t.Errorf("expected %d, got %d", n, el)
		}
		if err != nil {
			t.Error("queue error:", err)
		}
	}
}
