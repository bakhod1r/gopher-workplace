package fillptr

import "testing"

func TestMinMax(t *testing.T) {
	var lo, hi int
	MinMax([]int{3, -1, 7, 2}, &lo, &hi)
	if lo != -1 || hi != 7 {
		t.Errorf("lo,hi=%d,%d want -1,7", lo, hi)
	}
}
