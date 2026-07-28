package sumv

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(); got != 0 {
		t.Errorf("Sum()=%d want 0", got)
	}
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1,2,3)=%d want 6", got)
	}
	xs := []int{4, 5, 6}
	if got := Sum(xs...); got != 15 {
		t.Errorf("Sum(xs...)=%d want 15", got)
	}
}
