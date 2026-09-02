package maxslice

import "testing"

func TestMaxOf(t *testing.T) {
	if got, ok := MaxOf([]int{1, 9, 3}); got != 9 || !ok {
		t.Errorf("MaxOf([]int{1, 9, 3}) = %v, %v, want 9, true", got, ok)
	}
	if got, ok := MaxOf([]int{-5, -1}); got != -1 || !ok {
		t.Errorf("MaxOf([]int{-5, -1}) = %v, %v, want -1, true", got, ok)
	}
	if got, ok := MaxOf([]string{"a", "c", "b"}); got != "c" || !ok {
		t.Errorf("MaxOf([]string{...}) = %q, %v, want \"c\", true", got, ok)
	}
	if got, ok := MaxOf([]int{}); got != 0 || ok {
		t.Errorf("MaxOf([]int{}) = %v, %v, want 0, false", got, ok)
	}
}
