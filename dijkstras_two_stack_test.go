package dsa

import "testing"

func Test_evaluate(t *testing.T) {
	actual := evaluate("( 1 + ( ( 2 + 3 ) * ( 4 * 5 ) ) )")
	expected := 101
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
