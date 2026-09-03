package retptr

import "testing"

var sink *int

func TestNew(t *testing.T) {
	p := New(7)
	if p == nil || *p != 7 {
		t.Fatalf("New(7) = %v, want a pointer to 7", p)
	}
	if got := New(0); *got != 0 {
		t.Errorf("*New(0) = %d, want 0", *got)
	}
}

func TestNewReturnsDistinctPointers(t *testing.T) {
	a, b := New(1), New(1)
	if a == b {
		t.Error("New returned the same pointer twice")
	}
	*a = 99
	if *b != 1 {
		t.Errorf("*b = %d, want 1: the two ints share storage", *b)
	}
}

func TestNewAllocatesOnce(t *testing.T) {
	if n := testing.AllocsPerRun(100, func() { sink = New(3) }); n != 1 {
		t.Errorf("New made %v allocations, want exactly 1", n)
	}
}
