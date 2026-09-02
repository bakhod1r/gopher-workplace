package stackofstacksgen

import "testing"

func TestPlateStackSplits(t *testing.T) {
	s := PlateStack[int]{}
	s.Cap(2)
	s.Push(1)
	s.Push(2)
	if s.Stacks() != 1 {
		t.Errorf("Stacks() = %d, want 1", s.Stacks())
	}
	s.Push(3)
	if s.Stacks() != 2 {
		t.Errorf("Stacks() = %d, want 2", s.Stacks())
	}
	if v, ok := s.Pop(); v != 3 || !ok {
		t.Errorf("Pop() = %v, %v, want 3, true", v, ok)
	}
	if s.Stacks() != 1 {
		t.Errorf("Stacks() = %d, want 1 (an emptied inner stack must be dropped)", s.Stacks())
	}
}

func TestPlateStackOrder(t *testing.T) {
	s := PlateStack[int]{}
	s.Cap(2)
	for _, v := range []int{1, 2, 3, 4} {
		s.Push(v)
	}
	for _, w := range []int{4, 3, 2, 1} {
		got, ok := s.Pop()
		if !ok || got != w {
			t.Fatalf("Pop() = %v, %v, want %v, true", got, ok, w)
		}
	}
	if _, ok := s.Pop(); ok {
		t.Error("Pop() on empty reported ok")
	}
	if s.Stacks() != 0 {
		t.Errorf("Stacks() = %d, want 0", s.Stacks())
	}
}
