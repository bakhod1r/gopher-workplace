package thumbnailer

import (
	"reflect"
	"testing"
)

func TestTargetHeights(t *testing.T) {
	cases := []struct {
		name     string
		images   []Image
		maxWidth int
		want     []int
	}{
		{"landscape_and_square", []Image{{100, 50}, {200, 200}}, 100, []int{50, 100}},
		{"zero_width", []Image{{0, 40}}, 100, []int{0}},
		{"empty", []Image{}, 100, []int{}},
		{"upscale", []Image{{50, 20}}, 200, []int{80}},
		{"portrait", []Image{{40, 120}, {10, 10}}, 20, []int{60, 20}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TargetHeights(tc.images, tc.maxWidth); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TargetHeights(%v) = %v, want %v", tc.images, got, tc.want)
			}
		})
	}
}
