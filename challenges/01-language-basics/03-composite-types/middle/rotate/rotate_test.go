package rotate

import (
	"reflect"
	"testing"
)

func TestLeft(t *testing.T) {
	cases := []struct {
		xs   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{3, 4, 5, 1, 2}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 4, []int{2, 3, 1}},
		{[]int{1, 2, 3}, -1, []int{3, 1, 2}},
		{[]int{}, 2, []int{}},
	}
	for _, c := range cases {
		got := Left(c.xs, c.k)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Left(%v,%d)=%v; want %v", c.xs, c.k, got, c.want)
		}
	}
}
