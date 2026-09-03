package paddedcounters

import (
	"testing"
	"unsafe"
)

func TestRunTotals(t *testing.T) {
	if got := Run(4, 1000); got != 4000 {
		t.Errorf("Run = %d, want 4000", got)
	}
	if got := Run(1, 0); got != 0 {
		t.Errorf("Run = %d, want 0", got)
	}
	if got := Run(0, 10); got != 0 {
		t.Errorf("Run = %d, want 0", got)
	}
	if got := Run(3, -1); got != 0 {
		t.Errorf("Run = %d, want 0", got)
	}
}

func TestRunUnderLoad(t *testing.T) {
	if got := Run(8, 200000); got != 1600000 {
		t.Errorf("Run = %d, want 1600000", got)
	}
}

func TestSlotOccupiesAWholeLine(t *testing.T) {
	if got := unsafe.Sizeof(Slot{}); got != LineSize {
		t.Errorf("sizeof(Slot) = %d, want %d", got, LineSize)
	}
}

func TestSlotsAreALineApart(t *testing.T) {
	s := make([]Slot, 2)
	a := uintptr(unsafe.Pointer(&s[0]))
	b := uintptr(unsafe.Pointer(&s[1]))
	if b-a != LineSize {
		t.Errorf("stride = %d, want %d: neighbouring counters share a line", b-a, LineSize)
	}
}

func TestCounterFieldIsFirst(t *testing.T) {
	if off := unsafe.Offsetof(Slot{}.N); off != 0 {
		t.Errorf("N is at offset %d, want 0", off)
	}
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run(4, 20000)
	}
}
