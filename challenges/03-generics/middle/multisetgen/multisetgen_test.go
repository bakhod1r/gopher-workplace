package multisetgen

import "testing"

func TestBagCounts(t *testing.T) {
	b := NewBag[string]()
	if b.Count("a") != 0 {
		t.Errorf("Count on empty = %d, want 0", b.Count("a"))
	}
	b.Add("a")
	b.Add("a")
	b.Add("b")
	if b.Count("a") != 2 {
		t.Errorf("Count(a) = %d, want 2", b.Count("a"))
	}
	if b.Distinct() != 2 {
		t.Errorf("Distinct() = %d, want 2", b.Distinct())
	}
}

func TestBagRemove(t *testing.T) {
	b := NewBag[string]()
	b.Add("a")
	if !b.Remove("a") {
		t.Error(`Remove("a") = false, want true`)
	}
	if b.Distinct() != 0 {
		t.Errorf("Distinct() = %d, want 0 (the key must be deleted at zero)", b.Distinct())
	}
	if b.Remove("a") {
		t.Error("Remove on an absent value = true, want false")
	}
	if b.Count("a") != 0 {
		t.Errorf("Count(a) = %d, want 0", b.Count("a"))
	}
}

func TestBagRemovePartial(t *testing.T) {
	b := NewBag[int]()
	b.Add(1)
	b.Add(1)
	b.Remove(1)
	if b.Count(1) != 1 || b.Distinct() != 1 {
		t.Errorf("Count = %d, Distinct = %d, want 1, 1", b.Count(1), b.Distinct())
	}
}
