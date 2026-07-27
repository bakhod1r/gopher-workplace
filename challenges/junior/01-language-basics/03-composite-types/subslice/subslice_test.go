package subslice

import (
	"reflect"
	"testing"
)

func TestHead(t *testing.T) {
	cases := []struct {
		name string
		s    []int
		n    int
		want []int
	}{
		{"first two", []int{1, 2, 3, 4}, 2, []int{1, 2}},
		{"n over length clamps", []int{1, 2}, 5, []int{1, 2}},
		{"zero", []int{1, 2, 3}, 0, []int{}},
		{"negative", []int{1, 2, 3}, -3, []int{}},
		{"all", []int{7, 8}, 2, []int{7, 8}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Head(tc.s, tc.n); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Head(%v, %d) = %v, want %v", tc.s, tc.n, got, tc.want)
			}
		})
	}
}

func TestHeadIsIndependent(t *testing.T) {
	s := []int{1, 2, 3, 4}
	h := Head(s, 2)
	h[0] = 99
	if s[0] != 1 {
		t.Errorf("writing to result changed the input: s[0] = %d, want 1", s[0])
	}
	s[1] = 88
	if h[1] != 2 {
		t.Errorf("writing to input changed the result: h[1] = %d, want 2", h[1])
	}
}
