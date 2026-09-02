package comparegen

import "testing"

func TestCompare(t *testing.T) {
	if got := Compare(1, 2); got != -1 {
		t.Errorf("Compare(1, 2) = %v, want -1", got)
	}
	if got := Compare(2, 2); got != 0 {
		t.Errorf("Compare(2, 2) = %v, want 0", got)
	}
	if got := Compare(3, 2); got != 1 {
		t.Errorf("Compare(3, 2) = %v, want 1", got)
	}
	if got := Compare("b", "a"); got != 1 {
		t.Errorf(`Compare("b", "a") = %v, want 1`, got)
	}
	if got := Compare(1.5, 2.5); got != -1 {
		t.Errorf("Compare(1.5, 2.5) = %v, want -1", got)
	}
}
