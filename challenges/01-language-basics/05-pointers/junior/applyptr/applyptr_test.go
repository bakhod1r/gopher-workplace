package applyptr

import "testing"

func TestApply(t *testing.T) {
	x := 5
	Apply(&x, func(n int) int { return n * n })
	if x != 25 {
		t.Errorf("x=%d want 25", x)
	}
}
