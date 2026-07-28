package movezeros

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	cases := []struct{ xs, want []int }{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0}, []int{0, 0}},
		{nil, []int{}},
	}
	for _, c := range cases {
		got := MoveZeros(c.xs)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("MoveZeros(%v)=%v; want %v", c.xs, got, c.want)
		}
	}
}
