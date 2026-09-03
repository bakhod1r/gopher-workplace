package minstacksyncbug

import "testing"

func TestMinAfterPop(t *testing.T) {
	var s MinStack[int]
	s.Push(3)
	s.Push(1)
	if _, ok := s.Pop(); !ok {
		t.Fatal("Pop = false, want true")
	}
	got, ok := s.Min()
	if !ok || got != 3 {
		t.Errorf("Min = %d, %v, want 3, true", got, ok)
	}
}

func TestMinSingle(t *testing.T) {
	var s MinStack[int]
	s.Push(3)
	if got, ok := s.Min(); !ok || got != 3 {
		t.Errorf("Min = %d, %v, want 3, true", got, ok)
	}
}

func TestPopEmpty(t *testing.T) {
	var s MinStack[int]
	if got, ok := s.Pop(); ok || got != 0 {
		t.Errorf("Pop = %d, %v, want 0, false", got, ok)
	}
}
