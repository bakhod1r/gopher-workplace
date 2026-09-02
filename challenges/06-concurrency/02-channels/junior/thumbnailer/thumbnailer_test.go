package thumbnailer

import (
	"reflect"
	"testing"
)

func TestThumbAreas(t *testing.T) {
	cases := []struct {
		name  string
		sides []int
		want  []int
	}{
		{"two_sizes", []int{64, 128}, []int{4096, 16384}},
		{"tiny", []int{2}, []int{4}},
		{"empty", nil, []int{}},
		{"zero_side", []int{0, 4}, []int{0, 16}},
		{"order_kept", []int{3, 1, 2}, []int{9, 1, 4}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ThumbAreas(tc.sides)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ThumbAreas(%v) = %#v, want %#v", tc.sides, got, tc.want)
			}
		})
	}
}
