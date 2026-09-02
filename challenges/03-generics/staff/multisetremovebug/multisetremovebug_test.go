package multisetremovebug

import (
	"testing"
)

func TestRemoveDecrements(t *testing.T) {
	var s Multiset[string]
	for i := 0; i < 3; i++ {
		s.Add("x")
	}
	if !s.Remove("x") {
		t.Fatal("Remove(x) = false, want true")
	}
	if got := s.Count("x"); got != 2 {
		t.Errorf("Count(x) = %d, want 2", got)
	}
	if got := s.Len(); got != 2 {
		t.Errorf("Len = %d, want 2", got)
	}
	if got := s.Distinct(); got != 1 {
		t.Errorf("Distinct = %d, want 1", got)
	}
}

func TestRemoveLastOccurrence(t *testing.T) {
	var s Multiset[string]
	s.Add("x")
	if !s.Remove("x") {
		t.Fatal("Remove(x) = false, want true")
	}
	if s.Remove("x") {
		t.Error("second Remove(x) = true, want false")
	}
	if got := s.Distinct(); got != 0 {
		t.Errorf("Distinct = %d, want 0", got)
	}
	if got := s.Len(); got != 0 {
		t.Errorf("Len = %d, want 0", got)
	}
}

func TestRemoveMixed(t *testing.T) {
	var s Multiset[int]
	for _, v := range []int{1, 1, 1, 2, 2, 3} {
		s.Add(v)
	}
	s.Remove(1)
	s.Remove(2)
	if s.Len() != 4 {
		t.Errorf("Len = %d, want 4", s.Len())
	}
	if s.Count(1) != 2 || s.Count(2) != 1 || s.Count(3) != 1 {
		t.Errorf("counts = %d/%d/%d, want 2/1/1", s.Count(1), s.Count(2), s.Count(3))
	}
	if s.Distinct() != 3 {
		t.Errorf("Distinct = %d, want 3", s.Distinct())
	}
}
