package takewhile

import (
	"reflect"
	"testing"
)

func TestTakePositive(t *testing.T) {
	cases := []struct{ xs, want []int }{
		{[]int{1, 2, 3, -1, 4}, []int{1, 2, 3}},
		{[]int{-1, 2}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{nil, []int{}},
	}
	for _, c := range cases {
		got := TakePositive(c.xs)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("TakePositive(%v)=%v; want %v", c.xs, got, c.want)
		}
	}
}
