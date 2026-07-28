package sumints

import "testing"

func TestSum(t *testing.T) {
	cases := []struct {
		xs   []int
		want int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{}, 0},
		{nil, 0},
		{[]int{-5, 5}, 0},
	}
	for _, c := range cases {
		if got := Sum(c.xs); got != c.want {
			t.Errorf("Sum(%v)=%d; want %d", c.xs, got, c.want)
		}
	}
}
