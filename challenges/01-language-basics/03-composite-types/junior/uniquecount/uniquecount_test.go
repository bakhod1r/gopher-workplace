package uniquecount

import "testing"

func TestDistinct(t *testing.T) {
	cases := []struct {
		xs   []int
		want int
	}{
		{[]int{1, 2, 2, 3, 3, 3}, 3},
		{[]int{5, 5, 5}, 1},
		{nil, 0},
		{[]int{1, 2, 3}, 3},
	}
	for _, c := range cases {
		if got := Distinct(c.xs); got != c.want {
			t.Errorf("Distinct(%v)=%d; want %d", c.xs, got, c.want)
		}
	}
}
