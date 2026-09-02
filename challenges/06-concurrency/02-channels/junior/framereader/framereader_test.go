package framereader

import (
	"reflect"
	"testing"
)

func TestReadFrames(t *testing.T) {
	cases := []struct {
		name  string
		sizes []int
		extra int
		want  []int
	}{
		{"one_extra", []int{1024, 512}, 1, []int{1024, 512, 0}},
		{"stream_ended", nil, 2, []int{0, 0}},
		{"no_extra", []int{5}, 0, []int{5}},
		{"negative_extra", []int{5, 6}, -3, []int{5, 6}},
		{"mostly_extra", []int{7}, 3, []int{7, 0, 0, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ReadFrames(tc.sizes, tc.extra)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("ReadFrames(%v, %d) = %#v, want %#v",
					tc.sizes, tc.extra, got, tc.want)
			}
		})
	}
}
