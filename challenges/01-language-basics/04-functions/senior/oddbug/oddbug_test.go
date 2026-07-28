package oddbug

import "testing"

func TestCountOdd(t *testing.T) {
	if got := CountOdd([]int{1, 2, 3, -3, -4, -5}); got != 4 {
		t.Errorf("=%d want 4 (1,3,-3,-5)", got)
	}
}
