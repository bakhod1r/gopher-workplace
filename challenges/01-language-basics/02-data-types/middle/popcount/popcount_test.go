package popcount

import "testing"

func TestCount(t *testing.T) {
	cases := []struct {
		x    uint64
		want int
	}{
		{0, 0}, {1, 1}, {0b1011, 3}, {0xFF, 8}, {^uint64(0), 64},
	}
	for _, c := range cases {
		if got := Count(c.x); got != c.want {
			t.Errorf("Count(%#x)=%d; want %d", c.x, got, c.want)
		}
	}
}
