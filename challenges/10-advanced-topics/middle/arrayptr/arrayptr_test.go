package arrayptr

import "testing"

var sink int

func TestSum(t *testing.T) {
	a := [8]int{1, 2, 3}
	if got := Sum(&a); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
	if got := Sum(&[8]int{}); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %d, want 0", got)
	}
}

func TestSumSeesLaterWrites(t *testing.T) {
	a := [8]int{}
	p := &a
	a[0] = 5
	if got := Sum(p); got != 5 {
		t.Errorf("Sum = %d, want 5: the pointer must reach the caller's array", got)
	}
}

func TestSumAllocatesNothing(t *testing.T) {
	a := [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	if n := testing.AllocsPerRun(200, func() { sink = Sum(&a) }); n != 0 {
		t.Errorf("Sum made %v allocations, want 0", n)
	}
}

func TestSumNegative(t *testing.T) {
	a := [8]int{-4, 4}
	if got := Sum(&a); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
}
