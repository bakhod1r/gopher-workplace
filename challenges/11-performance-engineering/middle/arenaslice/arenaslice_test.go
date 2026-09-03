package arenaslice

import "testing"

var sink []byte

func TestAllocCarvesFromTheBlock(t *testing.T) {
	a := NewArena(64)
	b, ok := a.Alloc(16)
	if !ok || len(b) != 16 {
		t.Fatalf("Alloc = %v, %v, want 16 bytes and true", len(b), ok)
	}
	if a.Used() != 16 || a.Free() != 48 {
		t.Errorf("Used, Free = %d, %d, want 16, 48", a.Used(), a.Free())
	}
}

func TestAllocationsDoNotOverlap(t *testing.T) {
	a := NewArena(64)
	x, _ := a.Alloc(8)
	y, _ := a.Alloc(8)
	for i := range x {
		x[i] = 1
	}
	for _, v := range y {
		if v != 0 {
			t.Fatal("two allocations shared memory")
		}
	}
}

func TestAllocFailsWithoutConsuming(t *testing.T) {
	a := NewArena(16)
	if _, ok := a.Alloc(32); ok {
		t.Error("Alloc(32) from a 16-byte arena reported success")
	}
	if a.Used() != 0 {
		t.Errorf("Used = %d, want 0 — a failed allocation must not consume the block", a.Used())
	}
	if _, ok := a.Alloc(16); !ok {
		t.Error("Alloc(16) failed after a rejected larger request")
	}
}

func TestAllocZeroAndNegative(t *testing.T) {
	a := NewArena(16)
	for _, n := range []int{0, -8} {
		b, ok := a.Alloc(n)
		if !ok || b == nil || len(b) != 0 {
			t.Errorf("Alloc(%d) = %v, %v, want an empty non-nil slice and true", n, b, ok)
		}
	}
	if a.Used() != 0 {
		t.Errorf("Used = %d, want 0", a.Used())
	}
}

func TestResetFreesEverything(t *testing.T) {
	a := NewArena(32)
	a.Alloc(16)
	a.Alloc(16)
	if a.Free() != 0 {
		t.Fatalf("Free = %d, want 0", a.Free())
	}
	a.Reset()
	if a.Used() != 0 || a.Free() != 32 {
		t.Errorf("after Reset: Used, Free = %d, %d, want 0, 32", a.Used(), a.Free())
	}
}

func TestResetZeroesReusedMemory(t *testing.T) {
	a := NewArena(16)
	b, _ := a.Alloc(8)
	for i := range b {
		b[i] = 0xFF
	}
	a.Reset()
	again, _ := a.Alloc(8)
	for _, v := range again {
		if v != 0 {
			t.Fatal("recycled arena memory still holds the previous contents")
		}
	}
}

func TestAllocDoesNotAllocate(t *testing.T) {
	a := NewArena(1 << 20)
	allocs := testing.AllocsPerRun(100, func() {
		a.Reset()
		sink, _ = a.Alloc(4096)
	})
	if allocs != 0 {
		t.Errorf("Alloc made %v heap allocations, want 0 — the block is already there", allocs)
	}
}

func TestZeroSizedArena(t *testing.T) {
	a := NewArena(0)
	if _, ok := a.Alloc(1); ok {
		t.Error("Alloc from a zero-sized arena reported success")
	}
	if a.Free() != 0 {
		t.Errorf("Free = %d, want 0", a.Free())
	}
}
