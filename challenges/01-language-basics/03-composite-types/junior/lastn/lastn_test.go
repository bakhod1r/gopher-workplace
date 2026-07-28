package lastn

import (
	"reflect"
	"testing"
)

func TestLast(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	cases := []struct {
		n    int
		want []int
	}{
		{2, []int{4, 5}},
		{0, []int{}},
		{10, []int{1, 2, 3, 4, 5}},
		{-1, []int{}},
	}
	for _, c := range cases {
		got := Last(xs, c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Last(%v,%d)=%v; want %v", xs, c.n, got, c.want)
		}
	}
}
