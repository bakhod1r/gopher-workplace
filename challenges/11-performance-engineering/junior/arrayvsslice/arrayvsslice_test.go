package arrayvsslice

import "testing"

var sink int

func TestSumBlock(t *testing.T) {
	if got := SumBlock(Block{1, 2, 3}); got != 6 {
		t.Errorf("SumBlock = %d, want 6", got)
	}
	if got := SumBlock(Block{}); got != 0 {
		t.Errorf("SumBlock = %d, want 0", got)
	}
}

func TestZeroBlockLeavesTheArgumentAlone(t *testing.T) {
	b := Block{1, 2, 3}
	got := ZeroBlock(b)
	if got != (Block{}) {
		t.Errorf("ZeroBlock = %v, want the zero Block", got)
	}
	if b != (Block{1, 2, 3}) {
		t.Errorf("the argument changed: %v — arrays are passed by value", b)
	}
}

func TestSumSlice(t *testing.T) {
	if got := SumSlice([]int{1, 2, 3}); got != 6 {
		t.Errorf("SumSlice = %d, want 6", got)
	}
	if got := SumSlice(nil); got != 0 {
		t.Errorf("SumSlice(nil) = %d, want 0", got)
	}
}

func TestSlicingABlockSharesItsMemory(t *testing.T) {
	b := Block{1, 2, 3}
	s := b[:]
	s[0] = 99
	if b[0] != 99 {
		t.Errorf("b[0] = %d, want 99 — b[:] shares the array", b[0])
	}
	if got := SumBlock(b); got != SumSlice(s) {
		t.Errorf("SumBlock and SumSlice disagree: %d vs %d", SumBlock(b), SumSlice(s))
	}
}

func TestNeitherAllocates(t *testing.T) {
	b := Block{1, 2, 3, 4, 5, 6, 7, 8}
	if allocs := testing.AllocsPerRun(100, func() { sink = SumBlock(b) }); allocs != 0 {
		t.Errorf("SumBlock made %v allocations, want 0", allocs)
	}
	s := b[:]
	if allocs := testing.AllocsPerRun(100, func() { sink = SumSlice(s) }); allocs != 0 {
		t.Errorf("SumSlice made %v allocations, want 0", allocs)
	}
}
