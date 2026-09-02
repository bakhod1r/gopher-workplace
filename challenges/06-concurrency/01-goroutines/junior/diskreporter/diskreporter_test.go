package diskreporter

import (
	"reflect"
	"testing"
)

func TestLargestFiles(t *testing.T) {
	cases := []struct {
		name    string
		volumes [][]int
		want    []int
	}{
		{"two_volumes", [][]int{{30, 10}, {5}}, []int{30, 5}},
		{"empty_volume", [][]int{{}}, []int{0}},
		{"empty", [][]int{}, []int{}},
		{"last_is_largest", [][]int{{1, 2, 99}}, []int{99}},
		{"mixed", [][]int{{7}, {}, {4, 8}}, []int{7, 0, 8}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := LargestFiles(tc.volumes); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("LargestFiles(%v) = %v, want %v", tc.volumes, got, tc.want)
			}
		})
	}
}
