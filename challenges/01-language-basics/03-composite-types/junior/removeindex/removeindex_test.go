package removeindex

import (
	"reflect"
	"testing"
)

func TestRemoveAt(t *testing.T) {
	cases := []struct {
		xs   []int
		i    int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 1, []int{1, 3, 4}},
		{[]int{1, 2, 3}, 0, []int{2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
	}
	for _, c := range cases {
		got := RemoveAt(append([]int{}, c.xs...), c.i)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("RemoveAt(%v,%d)=%v; want %v", c.xs, c.i, got, c.want)
		}
	}
}
