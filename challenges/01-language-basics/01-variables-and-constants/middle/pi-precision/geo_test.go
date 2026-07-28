package geo

import (
	"math"
	"testing"
)

func TestPiPrecise(t *testing.T) {
	if math.Abs(float64(Pi)-math.Pi) > 1e-12 {
		t.Fatalf("Pi=%v not close to math.Pi", float64(Pi))
	}
}

func TestArea(t *testing.T) {
	cases := []struct{ r, want float64 }{
		{1, math.Pi},
		{2, 4 * math.Pi},
		{0, 0},
		{0.5, 0.25 * math.Pi},
	}
	for _, c := range cases {
		if got := Area(c.r); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("Area(%v)=%v; want %v", c.r, got, c.want)
		}
	}
}
