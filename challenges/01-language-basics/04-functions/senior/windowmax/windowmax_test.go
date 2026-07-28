package windowmax

import "testing"

func TestMaxWindow(t *testing.T) {
	if got := MaxWindow([]int{1, 2, 3, 4, 5}, 2); got != 9 {
		t.Errorf("=%d want 9 (4+5)", got)
	}
	if got := MaxWindow([]int{2, 1, 5, 1, 3, 2}, 3); got != 9 {
		t.Errorf("=%d want 9 (5+1+3)", got)
	}
}
