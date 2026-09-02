package setgen

import "testing"

func TestSet(t *testing.T) {
	s := NewSet[int]()
	if s.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", s.Len())
	}
	if s.Has(9) {
		t.Error("Has(9) on empty set = true, want false")
	}
	s.Add(1)
	s.Add(1)
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1 after adding 1 twice", s.Len())
	}
	if !s.Has(1) {
		t.Error("Has(1) = false, want true")
	}
	s.Add(2)
	if s.Len() != 2 {
		t.Errorf("Len() = %d, want 2", s.Len())
	}
}

func TestSetStrings(t *testing.T) {
	s := NewSet[string]()
	s.Add("a")
	if !s.Has("a") || s.Has("b") {
		t.Error("string set membership is wrong")
	}
}
