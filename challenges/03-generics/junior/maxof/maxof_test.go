package maxof

import "testing"

func TestMax(t *testing.T) {
	if got := Max(2, 5); got != 5 {
		t.Errorf("Max(2, 5) = %v, want 5", got)
	}
	if got := Max(5, 2); got != 5 {
		t.Errorf("Max(5, 2) = %v, want 5", got)
	}
	if got := Max("a", "b"); got != "b" {
		t.Errorf(`Max("a", "b") = %q, want "b"`, got)
	}
	if got := Max(3.5, 1.5); got != 3.5 {
		t.Errorf("Max(3.5, 1.5) = %v, want 3.5", got)
	}
	if got := Max(4, 4); got != 4 {
		t.Errorf("Max(4, 4) = %v, want 4", got)
	}
}
