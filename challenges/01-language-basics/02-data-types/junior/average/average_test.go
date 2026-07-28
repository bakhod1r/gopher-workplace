package average

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want float64
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{5}, 5},
		{"half", []int{1, 2}, 1.5},
		{"repeating decimal", []int{1, 2, 4}, 7.0 / 3.0},
		{"negatives", []int{-2, 2}, 0},
		{"mixed", []int{10, 20, 30, 5}, 16.25},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := Average(tc.in)
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("Average(%v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
