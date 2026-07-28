package gcdloop

import "testing"

func TestGCD(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{12, 8, 4}, {17, 5, 1}, {0, 9, 9}, {100, 10, 10},
	}
	for _, c := range cases {
		if got := GCD(c.a, c.b); got != c.want {
			t.Errorf("GCD(%d,%d)=%d want %d", c.a, c.b, got, c.want)
		}
	}
}
