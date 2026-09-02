package firstof

import "testing"

func TestFirst(t *testing.T) {
	if got, ok := First([]int{3, 1}); got != 3 || !ok {
		t.Errorf("First([]int{3, 1}) = %v, %v, want 3, true", got, ok)
	}
	if got, ok := First([]string{"a"}); got != "a" || !ok {
		t.Errorf("First([]string{\"a\"}) = %q, %v, want \"a\", true", got, ok)
	}
	if got, ok := First([]int{}); got != 0 || ok {
		t.Errorf("First([]int{}) = %v, %v, want 0, false", got, ok)
	}
	if got, ok := First([]string(nil)); got != "" || ok {
		t.Errorf("First(nil) = %q, %v, want \"\", false", got, ok)
	}
}
