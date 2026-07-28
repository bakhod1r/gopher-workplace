package shadow

import "testing"

func TestTally(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want int
	}{
		{"mixed signs", []int{1, -2, 3}, 4},
		{"nil slice", nil, 0},
		{"empty slice", []int{}, 0},
		{"all negative", []int{-1, -5}, 0},
		{"zeros ignored", []int{0, 0, 7}, 7},
		{"single positive", []int{5}, 5},
		{"accumulates across many", []int{1, 2, 3, 4, 5}, 15},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Tally(tc.in); got != tc.want {
				t.Errorf("Tally(%v) = %d, want %d", tc.in, got, tc.want)
			}
		})
	}
}
