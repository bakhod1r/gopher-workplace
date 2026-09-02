package defaultmapgen

import "testing"

func TestDefaultMap(t *testing.T) {
	m := NewDefaultMap[string](7)
	if got := m.Get("x"); got != 7 {
		t.Errorf(`Get("x") = %v, want 7`, got)
	}
	m.Put("x", 1)
	if got := m.Get("x"); got != 1 {
		t.Errorf(`Get("x") = %v, want 1`, got)
	}
	m.Put("y", 0)
	if got := m.Get("y"); got != 0 {
		t.Errorf(`Get("y") = %v, want 0 (a stored zero is not missing)`, got)
	}
}

func TestDefaultMapStrings(t *testing.T) {
	m := NewDefaultMap[int]("none")
	if got := m.Get(1); got != "none" {
		t.Errorf(`Get(1) = %q, want "none"`, got)
	}
	m.Put(1, "")
	if got := m.Get(1); got != "" {
		t.Errorf(`Get(1) = %q, want ""`, got)
	}
}
