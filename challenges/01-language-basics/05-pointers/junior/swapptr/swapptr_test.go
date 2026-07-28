package swapptr

import "testing"

func TestSwap(t *testing.T) {
	x, y := 1, 2
	Swap(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("x,y=%d,%d want 2,1", x, y)
	}
}
