package clockmod

import "testing"

func TestAddHours(t *testing.T) {
	cases := []struct {
		h, add, want int
	}{
		{10, 5, 15},
		{23, 1, 0},
		{0, -1, 23},
		{6, -30, 0},
		{12, 48, 12},
	}
	for _, c := range cases {
		if got := AddHours(c.h, c.add); got != c.want {
			t.Errorf("AddHours(%d,%d)=%d; want %d", c.h, c.add, got, c.want)
		}
	}
}
