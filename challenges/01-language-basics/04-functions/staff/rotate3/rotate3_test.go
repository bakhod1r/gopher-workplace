package rotate3

import "testing"

func TestRotateLeft(t *testing.T) {
	a, b, c := RotateLeft(1, 2, 3)
	if a != 2 || b != 3 || c != 1 {
		t.Errorf("=%d,%d,%d want 2,3,1", a, b, c)
	}
}
