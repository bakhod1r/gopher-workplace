package checkedadd

import (
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		a, b int64
		sum  int64
		ok   bool
	}{
		{1, 2, 3, true},
		{math.MaxInt64, 1, 0, false},
		{math.MaxInt64 - 1, 1, math.MaxInt64, true},
		{math.MinInt64, -1, 0, false},
		{-5, -5, -10, true},
	}
	for _, c := range cases {
		sum, ok := Add(c.a, c.b)
		if ok != c.ok || (ok && sum != c.sum) {
			t.Errorf("Add(%d,%d)=(%d,%v); want (%d,%v)", c.a, c.b, sum, ok, c.sum, c.ok)
		}
	}
}
