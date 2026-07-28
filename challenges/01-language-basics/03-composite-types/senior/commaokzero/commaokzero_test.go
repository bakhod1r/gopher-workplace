package commaokzero

import "testing"

func TestGetOr(t *testing.T) {
	m := map[string]int{"a": 5, "zero": 0}
	if got := GetOr(m, "a", 99); got != 5 {
		t.Errorf("present: got %d", got)
	}
	if got := GetOr(m, "zero", 99); got != 0 {
		t.Errorf("stored zero must win: got %d, want 0", got)
	}
	if got := GetOr(m, "missing", 99); got != 99 {
		t.Errorf("missing: got %d", got)
	}
}
