package coalescegen

import "testing"

func TestCoalesce(t *testing.T) {
	if got := Coalesce(0, 0, 5); got != 5 {
		t.Errorf("Coalesce(0, 0, 5) = %v, want 5", got)
	}
	if got := Coalesce(3, 5); got != 3 {
		t.Errorf("Coalesce(3, 5) = %v, want 3", got)
	}
	if got := Coalesce("", "a", "b"); got != "a" {
		t.Errorf(`Coalesce("", "a", "b") = %q, want "a"`, got)
	}
	if got := Coalesce(0, 0); got != 0 {
		t.Errorf("Coalesce(0, 0) = %v, want 0", got)
	}
	if got := Coalesce[string](); got != "" {
		t.Errorf(`Coalesce[string]() = %q, want ""`, got)
	}
}
