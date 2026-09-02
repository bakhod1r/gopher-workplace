package swaptypegen

import "testing"

func TestSwap(t *testing.T) {
	p := SamePair[int]{First: 1, Second: 2}
	p.Swap()
	if p.First != 2 || p.Second != 1 {
		t.Errorf("after Swap = %+v, want {2 1}", p)
	}
}

func TestOrdered(t *testing.T) {
	p := SamePair[int]{First: 2, Second: 1}
	lo, hi := p.Ordered()
	if lo != 1 || hi != 2 {
		t.Errorf("Ordered() = %v, %v, want 1, 2", lo, hi)
	}
	if p.First != 2 || p.Second != 1 {
		t.Errorf("Ordered mutated the receiver: %+v, want {2 1}", p)
	}
	q := SamePair[string]{First: "a", Second: "b"}
	slo, shi := q.Ordered()
	if slo != "a" || shi != "b" {
		t.Errorf("Ordered() = %q, %q, want a, b", slo, shi)
	}
}
