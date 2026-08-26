package maxval

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	cases := []struct {
		name string
		vals []float64
		want float64
	}{
		{"normal", []float64{3, 1, 2}, 3},
		{"single", []float64{42}, 42},
		{"negative", []float64{-5, -1, 0}, 0},
		{"empty", nil, math.Inf(-1)},
		{"all_same", []float64{7, 7, 7}, 7},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Stats{Values: tc.vals}
			got := s.Max()
			if got != tc.want {
				t.Errorf("Stats{%v}.Max() = %g, want %g", tc.vals, got, tc.want)
			}
		})
	}
}
