package copyarray

import "testing"

func TestBump(t *testing.T) {
	if got := Bump([4]int{1, 2, 3, 4}); got != [4]int{2, 3, 4, 5} {
		t.Errorf("Bump = %v, want [2 3 4 5]", got)
	}
	if got := Bump([4]int{}); got != [4]int{1, 1, 1, 1} {
		t.Errorf("Bump = %v, want [1 1 1 1]", got)
	}
	if got := Bump([4]int{-1, 0, 0, 0}); got != [4]int{0, 1, 1, 1} {
		t.Errorf("Bump = %v, want [0 1 1 1]", got)
	}
}

func TestBumpDoesNotTouchTheCaller(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	Bump(a)
	if a != [4]int{1, 2, 3, 4} {
		t.Errorf("a = %v, want [1 2 3 4]: the caller's array changed", a)
	}
}

func TestBumpAllocatesNothing(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	if n := testing.AllocsPerRun(100, func() { _ = Bump(a) }); n != 0 {
		t.Errorf("Bump made %v allocations, want 0: arrays are values", n)
	}
}
