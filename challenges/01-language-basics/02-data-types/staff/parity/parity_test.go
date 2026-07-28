package parity

import "testing"

func TestParity(t *testing.T) {
	cases := []struct {
		x    uint32
		want int
	}{
		{0, 0}, {1, 1}, {3, 0}, {7, 1}, {0xFF, 0}, {0x80000000, 1},
	}
	for _, c := range cases {
		if got := Parity(c.x); got != c.want {
			t.Errorf("Parity(%#x)=%d; want %d", c.x, got, c.want)
		}
	}
}
