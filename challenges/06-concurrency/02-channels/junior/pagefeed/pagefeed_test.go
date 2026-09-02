package pagefeed

import (
	"reflect"
	"testing"
)

func TestStreamPages(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []int
	}{
		{"three_pages", 3, []int{1, 2, 3}},
		{"one_page", 1, []int{1}},
		{"no_pages", 0, []int{}},
		{"negative", -5, []int{}},
		{"five_pages", 5, []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := StreamPages(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("StreamPages(%d) = %#v, want %#v", tc.n, got, tc.want)
			}
		})
	}
}
