package gcd

import "testing"

func TestGCD(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{12, 8, 4}, {17, 5, 1}, {0, 9, 9}, {9, 0, 9}, {0, 0, 0}, {-12, 8, 4},
	}
	for _, c := range cases {
		if got := GCD(c.a, c.b); got != c.want {
			t.Errorf("GCD(%d,%d)=%d; want %d", c.a, c.b, got, c.want)
		}
	}
}
