package max3

import "testing"

func TestMax3(t *testing.T) {
	cases := []struct{ a, b, c, want int }{
		{1, 2, 3, 3},
		{3, 2, 1, 3},
		{5, 9, 5, 9},
		{-1, -2, -3, -1},
	}
	for _, c := range cases {
		if got := Max3(c.a, c.b, c.c); got != c.want {
			t.Errorf("Max3(%d,%d,%d)=%d want %d", c.a, c.b, c.c, got, c.want)
		}
	}
}
