package setmethodgen

import "testing"

func TestUnion(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(2, 3)
	u := a.Union(b)
	if u.Len() != 3 {
		t.Errorf("Union.Len() = %d, want 3", u.Len())
	}
	if a.Len() != 2 || b.Len() != 2 {
		t.Errorf("operands mutated: a=%d b=%d", a.Len(), b.Len())
	}
}

func TestIntersect(t *testing.T) {
	a := NewSet(1, 2)
	b := NewSet(2, 3)
	i := a.Intersect(b)
	if i.Len() != 1 {
		t.Errorf("Intersect.Len() = %d, want 1", i.Len())
	}
	empty := NewSet(1).Intersect(NewSet(9))
	if empty.Len() != 0 {
		t.Errorf("Intersect.Len() = %d, want 0", empty.Len())
	}
}

func TestChaining(t *testing.T) {
	got := NewSet(1, 2, 3).Union(NewSet(4)).Intersect(NewSet(2, 4, 9))
	if got.Len() != 2 {
		t.Errorf("chained Len() = %d, want 2", got.Len())
	}
	if NewSet[string]().Len() != 0 {
		t.Error("empty set has a non-zero length")
	}
}
