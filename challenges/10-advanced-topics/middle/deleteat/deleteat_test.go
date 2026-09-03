package deleteat

import "testing"

func TestDeleteAt(t *testing.T) {
	a, b, c := &Item{ID: 1}, &Item{ID: 2}, &Item{ID: 3}
	got := DeleteAt([]*Item{a, b, c}, 1)
	if len(got) != 2 || got[0] != a || got[1] != c {
		t.Errorf("DeleteAt = %v, want [a c]", got)
	}
}

func TestDeleteAtOrderIsPreserved(t *testing.T) {
	s := make([]*Item, 6)
	for i := range s {
		s[i] = &Item{ID: i}
	}
	got := DeleteAt(s, 0)
	for i, p := range got {
		if p.ID != i+1 {
			t.Fatalf("got[%d].ID = %d, want %d", i, p.ID, i+1)
		}
	}
}

func TestDeleteAtClearsTheTail(t *testing.T) {
	s := []*Item{{ID: 1}, {ID: 2}, {ID: 3}}
	backing := s
	DeleteAt(s, 0)
	if backing[2] != nil {
		t.Error("the vacated slot still holds an item: it stays reachable")
	}
}

func TestDeleteAtOutOfRange(t *testing.T) {
	s := []*Item{{ID: 1}}
	for _, i := range []int{-1, 1, 99} {
		if got := DeleteAt(s, i); len(got) != 1 {
			t.Errorf("DeleteAt(_, %d) = %v, want the slice unchanged", i, got)
		}
	}
}

func TestDeleteAtLast(t *testing.T) {
	s := []*Item{{ID: 1}, {ID: 2}}
	got := DeleteAt(s, 1)
	if len(got) != 1 || got[0].ID != 1 {
		t.Errorf("DeleteAt = %v, want [1]", got)
	}
	if s[1] != nil {
		t.Error("the last slot was not cleared")
	}
}

func TestDeleteAtAllocatesNothing(t *testing.T) {
	s := make([]*Item, 64)
	for i := range s {
		s[i] = &Item{ID: i}
	}
	if n := testing.AllocsPerRun(100, func() { _ = DeleteAt(s[:64], 0) }); n != 0 {
		t.Errorf("DeleteAt made %v allocations, want 0", n)
	}
}
