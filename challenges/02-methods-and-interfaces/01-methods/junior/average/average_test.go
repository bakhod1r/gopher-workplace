package average

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	cases := []struct {
		name string
		vals []float64
		want float64
	}{
		{"three", []float64{2, 4, 6}, 4},
		{"empty", nil, 0},
		{"single", []float64{7}, 7},
		{"mixed", []float64{-2, 2}, 0},
		{"decimals", []float64{1.5, 2.5}, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Stats{Values: tc.vals}
			got := s.Average()
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Stats{%v}.Average() = %g, want %g", tc.vals, got, tc.want)
			}
		})
	}
}
