package round2

import (
	"math"
	"testing"
)

func TestRound(t *testing.T) {
	cases := []struct {
		x      float64
		places int
		want   float64
	}{
		{3.14159, 2, 3.14},
		{2.5, 0, 3},
		{-2.675, 2, -2.68},
		{1.005, 2, 1.0}, // float repr of 1.005 is slightly below
	}
	for _, c := range cases {
		got := Round(c.x, c.places)
		if math.Abs(got-c.want) > 1e-9 {
			t.Errorf("Round(%v,%d)=%v; want %v", c.x, c.places, got, c.want)
		}
	}
}
