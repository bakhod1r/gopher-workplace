package sparsegridgen

import "testing"

func TestGrid(t *testing.T) {
	g := NewGrid("-")
	if got := g.At(9, 9); got != "-" {
		t.Errorf("At(unset) = %q, want -", got)
	}
	g.Set(1, 1, "x")
	if got := g.At(1, 1); got != "x" {
		t.Errorf("At(1,1) = %q, want x", got)
	}
	if got := g.At(1, 2); got != "-" {
		t.Errorf("At(1,2) = %q, want -", got)
	}
	if g.Filled() != 1 {
		t.Errorf("Filled() = %d, want 1", g.Filled())
	}
}

func TestGridNegativeCoordinates(t *testing.T) {
	g := NewGrid(0)
	g.Set(-3, -4, 7)
	if got := g.At(-3, -4); got != 7 {
		t.Errorf("At(-3,-4) = %v, want 7", got)
	}
	if got := g.At(3, 4); got != 0 {
		t.Errorf("At(3,4) = %v, want 0", got)
	}
}

func TestGridDefaultStillCounts(t *testing.T) {
	g := NewGrid(0)
	g.Set(0, 0, 0)
	if g.Filled() != 1 {
		t.Errorf("Filled() = %d, want 1 (a cell set to the default is still set)", g.Filled())
	}
}
