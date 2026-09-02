package stackembed

import "testing"

func TestTracedStack(t *testing.T) {
	var s TracedStack[int]
	if _, ok := s.Last(); ok {
		t.Error("Last() before any push reported ok, want false")
	}
	s.Push(1)
	s.Push(2)
	if got := s.Pushes(); got != 2 {
		t.Errorf("Pushes() = %d, want 2", got)
	}
	if v, ok := s.Last(); v != 2 || !ok {
		t.Errorf("Last() = %v, %v, want 2, true", v, ok)
	}
	if got := s.Len(); got != 2 {
		t.Errorf("Len() = %d, want 2 (the embedded stack must receive the pushes)", got)
	}
}

func TestTracedStackStrings(t *testing.T) {
	var s TracedStack[string]
	s.Push("a")
	if v, ok := s.Last(); v != "a" || !ok {
		t.Errorf("Last() = %q, %v, want a, true", v, ok)
	}
	if s.Len() != 1 {
		t.Errorf("Len() = %d, want 1", s.Len())
	}
}
