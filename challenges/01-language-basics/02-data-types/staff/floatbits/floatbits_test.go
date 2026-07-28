package floatbits

import "testing"

func TestExponent(t *testing.T) {
	cases := []struct {
		x    float64
		want int
	}{
		{1, 0}, {2, 1}, {4, 2}, {0.5, -1}, {0.25, -2}, {3, 1}, {1024, 10},
	}
	for _, c := range cases {
		if got := Exponent(c.x); got != c.want {
			t.Errorf("Exponent(%v)=%d; want %d", c.x, got, c.want)
		}
	}
}
