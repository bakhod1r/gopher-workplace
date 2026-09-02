package stackgen

import "testing"

func TestStack(t *testing.T) {
	var s Stack[int]
	if s.Len() != 0 {
		t.Fatalf("Len() = %d, want 0", s.Len())
	}
	if _, ok := s.Pop(); ok {
		t.Error("Pop() on empty stack reported ok, want false")
	}
	s.Push(1)
	s.Push(2)
	if s.Len() != 2 {
		t.Errorf("Len() = %d, want 2", s.Len())
	}
	if v, ok := s.Pop(); v != 2 || !ok {
		t.Errorf("Pop() = %v, %v, want 2, true", v, ok)
	}
	if v, ok := s.Pop(); v != 1 || !ok {
		t.Errorf("Pop() = %v, %v, want 1, true", v, ok)
	}
	if s.Len() != 0 {
		t.Errorf("Len() = %d, want 0", s.Len())
	}
}

func TestStackStrings(t *testing.T) {
	var s Stack[string]
	s.Push("a")
	if v, ok := s.Pop(); v != "a" || !ok {
		t.Errorf("Pop() = %q, %v, want a, true", v, ok)
	}
}
