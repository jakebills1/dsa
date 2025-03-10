package dsa

import (
	"testing"
)

func TestArrayStack(t *testing.T) {
	s := newArrayStack[int]()
	s.push(1)
	s.push(2)
	s.push(3)
	testCases := []int{3, 2, 1}
	for i := 0; i < 3; i++ {
		val, _ := s.pop()
		if val != testCases[i] {
			t.Error("expected", testCases[i], "got", val)
		}
	}
}
