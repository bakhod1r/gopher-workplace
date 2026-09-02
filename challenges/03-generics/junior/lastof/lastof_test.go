package lastof

import "testing"

func TestLast(t *testing.T) {
	if got, ok := Last([]int{3, 1, 4}); got != 4 || !ok {
		t.Errorf("Last([]int{3, 1, 4}) = %v, %v, want 4, true", got, ok)
	}
	if got, ok := Last([]string{"a", "b"}); got != "b" || !ok {
		t.Errorf("Last([]string{\"a\", \"b\"}) = %q, %v, want \"b\", true", got, ok)
	}
	if got, ok := Last([]int{7}); got != 7 || !ok {
		t.Errorf("Last([]int{7}) = %v, %v, want 7, true", got, ok)
	}
	if got, ok := Last([]int{}); got != 0 || ok {
		t.Errorf("Last([]int{}) = %v, %v, want 0, false", got, ok)
	}
}
