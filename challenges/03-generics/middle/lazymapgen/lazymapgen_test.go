package lazymapgen

import "testing"

func TestZeroValueIsUsable(t *testing.T) {
	var s Store[string, int]
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}
	if v, ok := s.Get("a"); v != 0 || ok {
		t.Errorf("Get on a fresh store = %v, %v, want 0, false", v, ok)
	}
	s.Set("a", 1)
	if v, ok := s.Get("a"); v != 1 || !ok {
		t.Errorf("Get = %v, %v, want 1, true", v, ok)
	}
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1", s.Len())
	}
}

func TestSetTwice(t *testing.T) {
	var s Store[int, string]
	s.Set(1, "a")
	s.Set(1, "b")
	if v, _ := s.Get(1); v != "b" {
		t.Errorf("Get = %q, want b", v)
	}
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1", s.Len())
	}
}
