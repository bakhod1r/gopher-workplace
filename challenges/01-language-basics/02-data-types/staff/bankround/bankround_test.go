package bankround

import "testing"

func TestRound(t *testing.T) {
	cases := []struct{ x, want float64 }{
		{2.5, 2}, {3.5, 4}, {0.5, 0}, {1.5, 2}, {-2.5, -2}, {2.4, 2}, {2.6, 3},
	}
	for _, c := range cases {
		if got := Round(c.x); got != c.want {
			t.Errorf("Round(%v)=%v; want %v", c.x, got, c.want)
		}
	}
}
