package sumlocal

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3}); got != 6 {
		t.Errorf("Sum = %d, want 6", got)
	}
	if got := Sum(nil); got != 0 {
		t.Errorf("Sum(nil) = %d, want 0", got)
	}
	if got := Sum([]int{-3, 3}); got != 0 {
		t.Errorf("Sum = %d, want 0", got)
	}
}

func TestSumAllocatesNothing(t *testing.T) {
	s := make([]int, 512)
	for i := range s {
		s[i] = i
	}
	if n := testing.AllocsPerRun(100, func() { _ = Sum(s) }); n != 0 {
		t.Errorf("Sum made %v allocations, want 0", n)
	}
}
