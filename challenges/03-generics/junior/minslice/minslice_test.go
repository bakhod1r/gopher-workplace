package minslice

import "testing"

func TestMinOf(t *testing.T) {
	if got, ok := MinOf([]int{4, 1, 3}); got != 1 || !ok {
		t.Errorf("MinOf([]int{4, 1, 3}) = %v, %v, want 1, true", got, ok)
	}
	if got, ok := MinOf([]int{4, 7}); got != 4 || !ok {
		t.Errorf("MinOf([]int{4, 7}) = %v, %v, want 4, true (do not seed with zero)", got, ok)
	}
	if got, ok := MinOf([]string{"b", "a"}); got != "a" || !ok {
		t.Errorf("MinOf([]string{...}) = %q, %v, want \"a\", true", got, ok)
	}
	if got, ok := MinOf([]int{}); got != 0 || ok {
		t.Errorf("MinOf([]int{}) = %v, %v, want 0, false", got, ok)
	}
}
