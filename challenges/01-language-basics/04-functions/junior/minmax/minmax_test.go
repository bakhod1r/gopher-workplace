package minmax

import "testing"

func TestMinMax(t *testing.T) {
	cases := []struct {
		in     []int
		mn, mx int
	}{
		{[]int{3}, 3, 3},
		{[]int{1, 2, 3}, 1, 3},
		{[]int{5, -2, 9, 0}, -2, 9},
	}
	for _, c := range cases {
		mn, mx := MinMax(c.in)
		if mn != c.mn || mx != c.mx {
			t.Errorf("MinMax(%v)=%d,%d want %d,%d", c.in, mn, mx, c.mn, c.mx)
		}
	}
}
