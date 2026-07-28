package twoappend

import "testing"

func TestFork(t *testing.T) {
	base := make([]int, 1, 4) // spare capacity
	base[0] = 7
	x, last := Fork(base, 100, 200)
	if x[0] != 7 {
		t.Errorf("x[0]=%d want 7", x[0])
	}
	if last != 100 {
		t.Errorf("x last=%d want 100 (clobbered by second append)", last)
	}
}
