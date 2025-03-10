package dsa

import (
	"testing"
)

func TestLLStack(t *testing.T) {
	s := newLLStack()
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
