package sumrange

import "testing"

func TestSumRange(t *testing.T) {
	if SumRange(nil) != 0 {
		t.Errorf("nil should be 0")
	}
	if got := SumRange([]int{1, 2, 3, 4}); got != 10 {
		t.Errorf("=%d want 10", got)
	}
}
