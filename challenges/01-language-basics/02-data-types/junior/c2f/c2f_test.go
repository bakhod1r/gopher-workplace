package c2f

import (
	"math"
	"testing"
)

func TestToF(t *testing.T) {
	cases := []struct {
		c    int
		want float64
	}{
		{0, 32}, {100, 212}, {37, 98.6}, {-40, -40},
	}
	for _, tc := range cases {
		if got := ToF(tc.c); math.Abs(got-tc.want) > 1e-9 {
			t.Errorf("ToF(%d)=%v; want %v", tc.c, got, tc.want)
		}
	}
}
