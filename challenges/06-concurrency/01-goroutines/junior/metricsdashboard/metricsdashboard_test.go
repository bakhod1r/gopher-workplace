package metricsdashboard

import (
	"reflect"
	"testing"
)

func TestPeakPerWindow(t *testing.T) {
	cases := []struct {
		name    string
		samples []int
		window  int
		want    []int
	}{
		{"even_windows", []int{1, 9, 3, 4}, 2, []int{9, 4}},
		{"ragged_tail", []int{5, 2, 8}, 2, []int{5, 8}},
		{"one_per_window", []int{3, 1}, 1, []int{3, 1}},
		{"all_negative", []int{-4, -9}, 5, []int{-4}},
		{"bad_window", []int{1}, 0, []int(nil)},
		{"empty", []int{}, 3, []int{}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := PeakPerWindow(tc.samples, tc.window); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PeakPerWindow(%v) = %v, want %v", tc.samples, got, tc.want)
			}
		})
	}
}
