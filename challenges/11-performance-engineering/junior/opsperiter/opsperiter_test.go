package opsperiter

import "testing"

func TestOpsPerIter(t *testing.T) {
	cases := []struct{ total, iters, want int }{
		{10, 4, 3}, // 2.5 rounds up
		{9, 4, 2},  // 2.25 rounds down
		{8, 4, 2},  // exact
		{7, 2, 4},  // 3.5 rounds up
		{1, 3, 0},  // 0.33 rounds down
		{0, 5, 0},  // nothing done
		{100, 1, 100},
	}
	for _, c := range cases {
		if got := OpsPerIter(c.total, c.iters); got != c.want {
			t.Errorf("OpsPerIter(%d, %d) = %d, want %d", c.total, c.iters, got, c.want)
		}
	}
}

func TestOpsPerIterNonPositiveIters(t *testing.T) {
	for _, n := range []int{0, -1} {
		if got := OpsPerIter(100, n); got != 0 {
			t.Errorf("OpsPerIter(100, %d) = %d, want 0", n, got)
		}
	}
}
