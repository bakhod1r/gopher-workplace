package gotoretry

import "testing"

func TestSumUntil(t *testing.T) {
	s, u := SumUntil([]int{1, 2, 3, 4}, 5)
	if s != 6 || u != 3 {
		t.Errorf("=%d,%d want 6,3", s, u)
	}
	s, u = SumUntil([]int{1, 1}, 100)
	if s != 2 || u != 2 {
		t.Errorf("=%d,%d want 2,2", s, u)
	}
}
