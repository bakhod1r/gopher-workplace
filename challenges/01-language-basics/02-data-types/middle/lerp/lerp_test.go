package lerp

import (
	"math"
	"testing"
)

func TestLerp(t *testing.T) {
	cases := []struct{ a, b, tt, want float64 }{
		{0, 10, 0, 0}, {0, 10, 1, 10}, {0, 10, 0.5, 5}, {2, 4, 0.25, 2.5},
	}
	for _, c := range cases {
		if got := Lerp(c.a, c.b, c.tt); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("Lerp(%v,%v,%v)=%v; want %v", c.a, c.b, c.tt, got, c.want)
		}
	}
}
