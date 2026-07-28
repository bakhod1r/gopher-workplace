package minmax

import "testing"

func TestMinMax(t *testing.T) {
	cases := []struct {
		xs     []int
		mn, mx int
		ok     bool
	}{
		{[]int{3, 1, 4, 1, 5}, 1, 5, true},
		{[]int{7}, 7, 7, true},
		{[]int{-2, -9, -1}, -9, -1, true},
		{nil, 0, 0, false},
	}
	for _, c := range cases {
		mn, mx, ok := MinMax(c.xs)
		if mn != c.mn || mx != c.mx || ok != c.ok {
			t.Errorf("MinMax(%v)=(%d,%d,%v); want (%d,%d,%v)", c.xs, mn, mx, ok, c.mn, c.mx, c.ok)
		}
	}
}
