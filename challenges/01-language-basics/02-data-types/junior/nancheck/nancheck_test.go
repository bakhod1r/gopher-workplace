package nancheck

import (
	"math"
	"testing"
)

func TestFinite(t *testing.T) {
	cases := []struct {
		x    float64
		want bool
	}{
		{0, true},
		{3.14, true},
		{math.NaN(), false},
		{math.Inf(1), false},
		{math.Inf(-1), false},
	}
	for _, c := range cases {
		if got := Finite(c.x); got != c.want {
			t.Errorf("Finite(%v)=%v; want %v", c.x, got, c.want)
		}
	}
}
