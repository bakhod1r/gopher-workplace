package seedbug

import "testing"

func TestMinOfPositives(t *testing.T) {
	if v, ok := MinOf([]int{4, 7}); v != 4 || !ok {
		t.Errorf("MinOf = %v, %v, want 4, true", v, ok)
	}
	if v, ok := MinOf([]int{9}); v != 9 || !ok {
		t.Errorf("MinOf = %v, %v, want 9, true", v, ok)
	}
}

func TestMinOfNegatives(t *testing.T) {
	if v, ok := MinOf([]int{-5, -1}); v != -5 || !ok {
		t.Errorf("MinOf = %v, %v, want -5, true", v, ok)
	}
}

func TestMinOfStrings(t *testing.T) {
	if v, ok := MinOf([]string{"b", "c"}); v != "b" || !ok {
		t.Errorf("MinOf = %q, %v, want b, true", v, ok)
	}
}

func TestMinOfEmpty(t *testing.T) {
	if v, ok := MinOf([]int{}); v != 0 || ok {
		t.Errorf("MinOf(empty) = %v, %v, want 0, false", v, ok)
	}
}
