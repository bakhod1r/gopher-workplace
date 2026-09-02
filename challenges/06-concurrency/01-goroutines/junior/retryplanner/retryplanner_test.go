package retryplanner

import (
	"reflect"
	"testing"
)

func TestBackoffs(t *testing.T) {
	cases := []struct {
		name     string
		attempts []int
		baseMs   int
		want     []int
	}{
		{"doubling", []int{0, 1, 3}, 100, []int{100, 200, 800}},
		{"negative_attempt", []int{-1}, 100, []int{100}},
		{"empty", []int{}, 100, []int{}},
		{"other_base", []int{2}, 50, []int{200}},
		{"long_tail", []int{4, 5}, 10, []int{160, 320}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Backoffs(tc.attempts, tc.baseMs); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Backoffs(%v) = %v, want %v", tc.attempts, got, tc.want)
			}
		})
	}
}
