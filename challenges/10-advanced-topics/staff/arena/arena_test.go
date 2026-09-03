package arena

import (
	"testing"
	"unsafe"
)

func TestAllocSequential(t *testing.T) {
	a := NewArena(64)
	b1, ok := a.Alloc(8, 1)
	if !ok || len(b1) != 8 {
		t.Fatalf("Alloc = %d bytes, %v, want 8, true", len(b1), ok)
	}
	b2, ok := a.Alloc(4, 1)
	if !ok || len(b2) != 4 {
		t.Fatalf("Alloc = %d bytes, %v, want 4, true", len(b2), ok)
	}
	b1[0] = 1
	b2[0] = 2
	if b1[0] != 1 || b2[0] != 2 {
		t.Error("the blocks overlap")
	}
}

func TestAllocAligns(t *testing.T) {
	a := NewArena(64)
	a.Alloc(1, 1)
	b, ok := a.Alloc(8, 8)
	if !ok {
		t.Fatal("Alloc reported false")
	}
	if uintptr(unsafe.Pointer(&b[0]))%8 != 0 {
		t.Error("the block is not 8-byte aligned")
	}
	if a.Used() < 16 {
		t.Errorf("Used = %d, want at least 16: the padding counts", a.Used())
	}
}

func TestAllocCapacityIsExact(t *testing.T) {
	a := NewArena(64)
	b, _ := a.Alloc(8, 1)
	if cap(b) != 8 {
		t.Fatalf("cap = %d, want 8", cap(b))
	}
	b = append(b, 'x')
	c, _ := a.Alloc(8, 1)
	if c[0] == 'x' {
		t.Error("appending to one block wrote into the next")
	}
}

func TestAllocRefusesWhenFull(t *testing.T) {
	a := NewArena(16)
	if _, ok := a.Alloc(16, 1); !ok {
		t.Fatal("a request for the whole arena must succeed")
	}
	if _, ok := a.Alloc(1, 1); ok {
		t.Error("Alloc reported true for a full arena")
	}
}

func TestAllocRejectsBadArguments(t *testing.T) {
	a := NewArena(64)
	for _, c := range []struct {
		n     int
		align uintptr
	}{
		{-1, 1}, {8, 0}, {8, 3}, {8, 6}, {65, 1},
	} {
		if _, ok := a.Alloc(c.n, c.align); ok {
			t.Errorf("Alloc(%d, %d) reported true, want false", c.n, c.align)
		}
	}
	if a.Used() != 0 {
		t.Errorf("Used = %d, want 0: a rejected request must consume nothing", a.Used())
	}
}

func TestAllocZeroBytes(t *testing.T) {
	a := NewArena(8)
	b, ok := a.Alloc(0, 1)
	if !ok || len(b) != 0 {
		t.Errorf("Alloc(0) = %d bytes, %v, want 0, true", len(b), ok)
	}
}
