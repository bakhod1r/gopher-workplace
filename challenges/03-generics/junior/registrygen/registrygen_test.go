package registrygen

import "testing"

func TestRegistry(t *testing.T) {
	r := NewRegistry[string, int]()
	if r.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", r.Len())
	}
	if !r.Register("a", 1) {
		t.Error(`Register("a", 1) = false, want true`)
	}
	if r.Register("a", 2) {
		t.Error(`Register("a", 2) = true, want false (duplicate)`)
	}
	if v, ok := r.Lookup("a"); v != 1 || !ok {
		t.Errorf(`Lookup("a") = %v, %v, want 1, true (first write wins)`, v, ok)
	}
	if _, ok := r.Lookup("missing"); ok {
		t.Error(`Lookup("missing") reported ok, want false`)
	}
	if r.Len() != 1 {
		t.Errorf("Len() = %d, want 1", r.Len())
	}
}
