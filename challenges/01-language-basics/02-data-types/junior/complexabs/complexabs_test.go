package complexabs

import (
	"math"
	"testing"
)

func TestMagnitude(t *testing.T) {
	cases := []struct {
		c    complex128
		want float64
	}{
		{complex(3, 4), 5},
		{complex(0, 0), 0},
		{complex(-3, -4), 5},
		{complex(1, 0), 1},
	}
	for _, c := range cases {
		if got := Magnitude(c.c); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("Magnitude(%v)=%v; want %v", c.c, got, c.want)
		}
	}
}
