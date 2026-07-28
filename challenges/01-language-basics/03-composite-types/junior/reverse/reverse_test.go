package reverse

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"odd length", []int{1, 2, 3}, []int{3, 2, 1}},
		{"even length", []int{1, 2}, []int{2, 1}},
		{"single", []int{9}, []int{9}},
		{"empty", []int{}, []int{}},
		{"longer", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			Reverse(tc.in)
			if !reflect.DeepEqual(tc.in, tc.want) {
				t.Errorf("after Reverse, got %v, want %v", tc.in, tc.want)
			}
		})
	}
}
