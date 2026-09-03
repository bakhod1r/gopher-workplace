package mapget

import "testing"

func TestGet(t *testing.T) {
	m := map[string]int{"a": 1, "zero": 0}
	if v, ok := Get(m, "a"); !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v, want 1, true", v, ok)
	}
	if v, ok := Get(m, "zero"); !ok || v != 0 {
		t.Errorf("Get(zero) = %d, %v, want 0, true: a stored zero is present", v, ok)
	}
	if v, ok := Get(m, "missing"); ok || v != 0 {
		t.Errorf("Get(missing) = %d, %v, want 0, false", v, ok)
	}
}

func TestGetNilMap(t *testing.T) {
	if v, ok := Get(nil, "a"); ok || v != 0 {
		t.Errorf("Get(nil) = %d, %v, want 0, false", v, ok)
	}
}

func TestGetEmptyKey(t *testing.T) {
	m := map[string]int{"": 5}
	if v, ok := Get(m, ""); !ok || v != 5 {
		t.Errorf("Get(\"\") = %d, %v, want 5, true", v, ok)
	}
}
