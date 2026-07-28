package clamp01

import (
	"math"
	"testing"
)

func TestSaturate(t *testing.T) {
	cases := []struct{ x, want float64 }{
		{0.5, 0.5}, {-1, 0}, {2, 1}, {0, 0}, {1, 1}, {math.NaN(), 0},
	}
	for _, c := range cases {
		if got := Saturate(c.x); got != c.want {
			t.Errorf("Saturate(%v)=%v; want %v", c.x, got, c.want)
		}
	}
}
