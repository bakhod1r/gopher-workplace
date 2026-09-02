package minstackgen

import "testing"

func TestMinStack(t *testing.T) {
	var s MinStack[int]
	if _, ok := s.Min(); ok {
		t.Error("Min() on empty reported ok")
	}
	s.Push(3)
	if v, ok := s.Min(); v != 3 || !ok {
		t.Errorf("Min() = %v, %v, want 3, true", v, ok)
	}
	s.Push(1)
	if v, _ := s.Min(); v != 1 {
		t.Errorf("Min() = %v, want 1", v)
	}
	s.Push(2)
	if v, _ := s.Min(); v != 1 {
		t.Errorf("Min() = %v, want 1", v)
	}
}

func TestMinStackPopRestoresMin(t *testing.T) {
	var s MinStack[int]
	s.Push(3)
	s.Push(1)
	if v, ok := s.Pop(); v != 1 || !ok {
		t.Fatalf("Pop() = %v, %v, want 1, true", v, ok)
	}
	if v, _ := s.Min(); v != 3 {
		t.Errorf("Min() after pop = %v, want 3", v)
	}
	s.Pop()
	if _, ok := s.Min(); ok {
		t.Error("Min() on drained stack reported ok")
	}
	if _, ok := s.Pop(); ok {
		t.Error("Pop() on empty reported ok")
	}
}

func TestMinStackDuplicateMins(t *testing.T) {
	var s MinStack[int]
	s.Push(1)
	s.Push(1)
	s.Pop()
	if v, ok := s.Min(); v != 1 || !ok {
		t.Errorf("Min() = %v, %v, want 1, true", v, ok)
	}
}
