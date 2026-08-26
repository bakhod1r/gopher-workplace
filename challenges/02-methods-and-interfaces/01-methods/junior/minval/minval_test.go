package minval

import (
	"math"
	"testing"
)

func TestMin(t *testing.T) {
	cases := []struct {
		name string
		vals []float64
		want float64
	}{
		{"normal", []float64{3, 1, 2}, 1},
		{"single", []float64{42}, 42},
		{"negative", []float64{-5, -1, 0}, -5},
		{"empty", nil, math.Inf(1)},
		{"all_same", []float64{7, 7, 7}, 7},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			s := Stats{Values: tc.vals}
			got := s.Min()
			if got != tc.want {
				t.Errorf("Stats{%v}.Min() = %g, want %g", tc.vals, got, tc.want)
			}
		})
	}
}
