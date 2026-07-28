package filterpos

import (
	"reflect"
	"testing"
)

func TestPositives(t *testing.T) {
	cases := []struct {
		xs   []int
		want []int
	}{
		{[]int{1, -2, 3, 0, 4}, []int{1, 3, 4}},
		{[]int{-1, -2}, []int{}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		got := Positives(c.xs)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Positives(%v)=%v; want %v", c.xs, got, c.want)
		}
	}
}
