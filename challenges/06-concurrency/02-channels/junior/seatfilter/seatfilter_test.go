package seatfilter

import (
	"reflect"
	"testing"
)

func TestAisleSeats(t *testing.T) {
	cases := []struct {
		name  string
		seats []int
		want  []int
	}{
		{"mixed_row", []int{1, 2, 3, 4}, []int{2, 4}},
		{"no_aisle_seats", []int{1, 3}, []int{}},
		{"empty_row", nil, []int{}},
		{"all_aisle", []int{2, 4, 6}, []int{2, 4, 6}},
		{"negatives_and_zero", []int{-4, -3, 0}, []int{-4, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := AisleSeats(tc.seats)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("AisleSeats(%v) = %#v, want %#v", tc.seats, got, tc.want)
			}
		})
	}
}
