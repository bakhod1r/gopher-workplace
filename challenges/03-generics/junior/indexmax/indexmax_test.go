package indexmax

import "testing"

func TestIndexOfMax(t *testing.T) {
	if got := IndexOfMax([]int{1, 9, 9}); got != 1 {
		t.Errorf("IndexOfMax([]int{1, 9, 9}) = %v, want 1 (first maximum)", got)
	}
	if got := IndexOfMax([]string{"c", "a"}); got != 0 {
		t.Errorf(`IndexOfMax([]string{"c", "a"}) = %v, want 0`, got)
	}
	if got := IndexOfMax([]int{-5, -1, -3}); got != 1 {
		t.Errorf("IndexOfMax([]int{-5, -1, -3}) = %v, want 1", got)
	}
	if got := IndexOfMax([]int{}); got != -1 {
		t.Errorf("IndexOfMax([]int{}) = %v, want -1", got)
	}
}
