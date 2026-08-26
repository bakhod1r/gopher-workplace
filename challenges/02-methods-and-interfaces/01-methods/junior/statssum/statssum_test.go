package statssum

import (
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name string
		vals []float64
		want float64
	}{
		{"three", []float64{1, 2, 3}, 6},
		{"empty", nil, 0},
		{"single", []float64{42}, 42},
		{"negatives", []float64{-1, -2, 3}, 0},
		{"decimals", []float64{0.1, 0.2, 0.3}, 0.6},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Stats{Values: tc.vals}
			got := s.Sum()
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Stats{%v}.Sum() = %g, want %g", tc.vals, got, tc.want)
			}
		})
	}
}
