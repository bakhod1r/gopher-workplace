package gridsum

import "testing"

func TestGridSum(t *testing.T) {
	if GridSum(nil) != 0 {
		t.Errorf("nil grid should be 0")
	}
	g := [][]int{{1, 2}, {3, 4}, {5}}
	if got := GridSum(g); got != 15 {
		t.Errorf("=%d want 15", got)
	}
}
