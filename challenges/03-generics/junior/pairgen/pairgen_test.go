package pairgen

import "testing"

func TestMakePair(t *testing.T) {
	p := MakePair(1, "a")
	if p.First != 1 || p.Second != "a" {
		t.Errorf("MakePair(1, %q) = %+v, want {1 a}", "a", p)
	}
}

func TestSwapped(t *testing.T) {
	s := MakePair(1, "a").Swapped()
	if s.First != "a" || s.Second != 1 {
		t.Errorf("Swapped() = %+v, want {a 1}", s)
	}
	b := MakePair(true, 2).Swapped()
	if b.First != 2 || b.Second != true {
		t.Errorf("Swapped() = %+v, want {2 true}", b)
	}
	back := MakePair(1, "a").Swapped().Swapped()
	if back.First != 1 || back.Second != "a" {
		t.Errorf("double Swapped() = %+v, want {1 a}", back)
	}
}
