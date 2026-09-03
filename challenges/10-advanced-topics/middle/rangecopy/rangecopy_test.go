package rangecopy

import "testing"

func TestBump(t *testing.T) {
	items := []Counter{{N: 1}, {N: 2}, {N: 3}}
	Bump(items)
	for i, c := range items {
		if c.N != i+2 {
			t.Errorf("items[%d].N = %d, want %d: the loop wrote to a copy", i, c.N, i+2)
		}
	}
}

func TestBumpEmpty(t *testing.T) {
	Bump(nil)
	Bump([]Counter{})
}

func TestBumpWritesThroughAView(t *testing.T) {
	items := []Counter{{N: 0}, {N: 0}, {N: 0}}
	Bump(items[1:2])
	if items[0].N != 0 || items[2].N != 0 {
		t.Errorf("items = %v, want only the middle element bumped", []int{items[0].N, items[1].N, items[2].N})
	}
	if items[1].N != 1 {
		t.Errorf("items[1].N = %d, want 1", items[1].N)
	}
}

func TestBumpAllocatesNothing(t *testing.T) {
	items := make([]Counter, 64)
	if n := testing.AllocsPerRun(100, func() { Bump(items) }); n != 0 {
		t.Errorf("Bump made %v allocations, want 0", n)
	}
}
