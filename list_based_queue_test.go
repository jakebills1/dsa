package dsa

import "testing"

func Test_ListBasedQueue(t *testing.T) {
	subject := newListBasedQueue()
	testCases := []int{1, 2, 3}
	for _, n := range testCases {
		subject.enqueue(n)
	}
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
