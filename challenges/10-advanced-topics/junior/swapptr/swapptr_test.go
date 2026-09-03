package swapptr

import "testing"

func TestSwap(t *testing.T) {
	x, y := 1, 2
	Swap(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("x, y = %d, %d, want 2, 1", x, y)
	}
}

func TestSwapSamePointer(t *testing.T) {
	x := 5
	Swap(&x, &x)
	if x != 5 {
		t.Errorf("x = %d, want 5", x)
	}
}

func TestSwapNil(t *testing.T) {
	x := 1
	Swap(&x, nil)
	Swap(nil, &x)
	Swap(nil, nil)
	if x != 1 {
		t.Errorf("x = %d, want 1", x)
	}
}

func TestSwapAllocatesNothing(t *testing.T) {
	x, y := 1, 2
	if n := testing.AllocsPerRun(200, func() { Swap(&x, &y) }); n != 0 {
		t.Errorf("Swap made %v allocations, want 0", n)
	}
}
