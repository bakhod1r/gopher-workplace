package accumulate

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		xs   []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{-1, 5, -2, 4}, 9},
		{[]int{}, 0},
		{[]int{-3, -4}, 0},
	}
	for _, c := range cases {
		if got := SumPositive(c.xs); got != c.want {
			t.Errorf("SumPositive(%v)=%d; want %d", c.xs, got, c.want)
		}
	}
}
