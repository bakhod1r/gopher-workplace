package feepipeline

import (
	"reflect"
	"testing"
)

func TestLineTotals(t *testing.T) {
	cases := []struct {
		name  string
		units []int
		want  []int
	}{
		{"two_lines", []int{1, 2}, []int{3, 5}},
		{"zero_units", []int{0}, []int{1}},
		{"empty_order", nil, []int{}},
		{"refund_lines", []int{-1, -2}, []int{-1, -3}},
		{"order_kept", []int{5, 3, 4}, []int{11, 7, 9}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := LineTotals(tc.units)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LineTotals(%v) = %#v, want %#v", tc.units, got, tc.want)
			}
		})
	}
}
