package storagereporter

import (
	"reflect"
	"testing"
)

func TestSectionKilobytes(t *testing.T) {
	cases := []struct {
		name     string
		sections [][]int
		want     []int
	}{
		{"two_sections", [][]int{{2048, 1024}, {4096}}, []int{3, 4}},
		{"empty_section", [][]int{{}}, []int{0}},
		{"empty", [][]int{}, []int{}},
		{"rounds_down", [][]int{{1023, 2047}}, []int{1}},
		{"mixed", [][]int{{1024}, {}, {1024, 1024, 1024}}, []int{1, 0, 3}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := SectionKilobytes(tc.sections); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SectionKilobytes(%v) = %v, want %v", tc.sections, got, tc.want)
			}
		})
	}
}
