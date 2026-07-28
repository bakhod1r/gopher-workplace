package aliasbug

import "testing"

func TestWithFirst(t *testing.T) {
	xs := []int{1, 2, 3}
	got := WithFirst(xs, 9)
	if got[0] != 9 {
		t.Errorf("result first=%d want 9", got[0])
	}
	if xs[0] != 1 {
		t.Errorf("input mutated: %v", xs)
	}
}
