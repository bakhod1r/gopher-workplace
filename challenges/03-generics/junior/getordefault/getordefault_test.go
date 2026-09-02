package getordefault

import "testing"

func TestGetOr(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}
	if got := GetOr(m, "a", 9); got != 1 {
		t.Errorf(`GetOr(m, "a", 9) = %v, want 1`, got)
	}
	if got := GetOr(m, "zero", 9); got != 0 {
		t.Errorf(`GetOr(m, "zero", 9) = %v, want 0 (stored zero, not the default)`, got)
	}
	if got := GetOr(m, "missing", 9); got != 9 {
		t.Errorf(`GetOr(m, "missing", 9) = %v, want 9`, got)
	}
	if got := GetOr(map[int]string(nil), 1, "def"); got != "def" {
		t.Errorf(`GetOr(nil, 1, "def") = %q, want "def"`, got)
	}
}
