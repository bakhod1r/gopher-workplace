package dedupesorted

import (
	"reflect"
	"testing"
)

func TestDedupe(t *testing.T) {
	cases := []struct {
		xs   []int
		want []int
	}{
		{[]int{1, 1, 2, 3, 3, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{5, 5, 5}, []int{5}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		got := Dedupe(c.xs)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Dedupe(%v)=%v; want %v", c.xs, got, c.want)
		}
	}
}
