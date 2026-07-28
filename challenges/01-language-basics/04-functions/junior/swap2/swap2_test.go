package swap2

import "testing"

func TestSwap(t *testing.T) {
	x, y := Swap(1, 2)
	if x != 2 || y != 1 {
		t.Errorf("Swap(1,2)=%d,%d want 2,1", x, y)
	}
	x, y = Swap(-5, 5)
	if x != 5 || y != -5 {
		t.Errorf("Swap(-5,5)=%d,%d want 5,-5", x, y)
	}
}
