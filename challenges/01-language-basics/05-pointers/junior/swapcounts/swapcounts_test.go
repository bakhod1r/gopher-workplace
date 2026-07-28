package swapcounts

import "testing"

func TestSwapCounts(t *testing.T) {
	x := Cart{Count: 1}
	y := Cart{Count: 9}
	SwapCounts(&x, &y)
	if x.Count != 9 || y.Count != 1 {
		t.Errorf("x,y=%d,%d want 9,1", x.Count, y.Count)
	}
}
