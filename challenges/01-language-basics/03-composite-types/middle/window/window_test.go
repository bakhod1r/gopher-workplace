package window

import (
	"reflect"
	"testing"
)

func TestSums(t *testing.T) {
	cases := []struct {
		xs   []int
		k    int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 2, []int{3, 5, 7}},
		{[]int{1, 2, 3}, 3, []int{6}},
		{[]int{1, 2, 3}, 5, []int{}},
		{[]int{5}, 1, []int{5}},
	}
	for _, c := range cases {
		got := Sums(c.xs, c.k)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Sums(%v,%d)=%v; want %v", c.xs, c.k, got, c.want)
		}
	}
}
