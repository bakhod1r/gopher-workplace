package mergesorted

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	cases := []struct{ a, b, want []int }{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 1, 2}, []int{1, 3}, []int{1, 1, 1, 2, 3}},
		{[]int{}, []int{1, 2}, []int{1, 2}},
		{nil, nil, []int{}},
	}
	for _, c := range cases {
		got := Merge(c.a, c.b)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Merge(%v,%v)=%v; want %v", c.a, c.b, got, c.want)
		}
	}
}
