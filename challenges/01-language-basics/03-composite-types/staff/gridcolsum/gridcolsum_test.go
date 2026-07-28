package gridcolsum

import "testing"

func TestColSum(t *testing.T) {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	cases := []struct{ c, want int }{
		{0, 12}, // 1+4+7
		{1, 15}, // 2+5+8
		{2, 18}, // 3+6+9
	}
	for _, tc := range cases {
		if got := ColSum(grid, tc.c); got != tc.want {
			t.Errorf("ColSum(c=%d)=%d; want %d", tc.c, got, tc.want)
		}
	}
}
