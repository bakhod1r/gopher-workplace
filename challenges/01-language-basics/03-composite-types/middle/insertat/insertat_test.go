package insertat

import (
	"reflect"
	"testing"
)

func TestInsertAt(t *testing.T) {
	cases := []struct {
		xs   []int
		i, v int
		want []int
	}{
		{[]int{1, 2, 3}, 1, 9, []int{1, 9, 2, 3}},
		{[]int{1, 2, 3}, 0, 9, []int{9, 1, 2, 3}},
		{[]int{1, 2, 3}, 3, 9, []int{1, 2, 3, 9}},
		{[]int{1, 2, 3}, 10, 9, []int{1, 2, 3, 9}},
		{[]int{}, 0, 9, []int{9}},
	}
	for _, c := range cases {
		got := InsertAt(append([]int{}, c.xs...), c.i, c.v)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("InsertAt(%v,%d,%d)=%v; want %v", c.xs, c.i, c.v, got, c.want)
		}
	}
}
