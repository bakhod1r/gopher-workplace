package minof

import "testing"

func TestMin(t *testing.T) {
	if got := Min(2, 5); got != 2 {
		t.Errorf("Min(2, 5) = %v, want 2", got)
	}
	if got := Min(5, 2); got != 2 {
		t.Errorf("Min(5, 2) = %v, want 2", got)
	}
	if got := Min("b", "a"); got != "a" {
		t.Errorf(`Min("b", "a") = %q, want "a"`, got)
	}
	if got := Min(3.5, 1.5); got != 1.5 {
		t.Errorf("Min(3.5, 1.5) = %v, want 1.5", got)
	}
	if got := Min(4, 4); got != 4 {
		t.Errorf("Min(4, 4) = %v, want 4", got)
	}
}
