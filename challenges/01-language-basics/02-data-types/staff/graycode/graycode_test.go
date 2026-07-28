package graycode

import "testing"

func TestToGray(t *testing.T) {
	// Successive Gray codes differ by exactly one bit.
	cases := []struct {
		x, want uint32
	}{
		{0, 0}, {1, 1}, {2, 3}, {3, 2}, {4, 6}, {5, 7}, {6, 5}, {7, 4},
	}
	for _, c := range cases {
		if got := ToGray(c.x); got != c.want {
			t.Errorf("ToGray(%d)=%d; want %d", c.x, got, c.want)
		}
	}
}
