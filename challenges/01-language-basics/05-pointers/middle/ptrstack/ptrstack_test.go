package ptrstack

import "testing"

func TestStack(t *testing.T) {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.Len() != 3 {
		t.Fatalf("len=%d want 3", s.Len())
	}
	if v, ok := s.Pop(); !ok || v != 3 {
		t.Fatalf("pop=%d,%v want 3,true", v, ok)
	}
}
