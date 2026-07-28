package epsiloneq

import (
	"math"
	"testing"
)

func TestEqual(t *testing.T) {
	cases := []struct {
		a, b, eps float64
		want      bool
	}{
		{0.1 + 0.2, 0.3, 1e-9, true},
		{1.0, 1.001, 1e-9, false},
		{1.0, 1.0, 0, true},
		{math.NaN(), math.NaN(), 1, false},
	}
	for _, c := range cases {
		if got := Equal(c.a, c.b, c.eps); got != c.want {
			t.Errorf("Equal(%v,%v,%v)=%v; want %v", c.a, c.b, c.eps, got, c.want)
		}
	}
}
