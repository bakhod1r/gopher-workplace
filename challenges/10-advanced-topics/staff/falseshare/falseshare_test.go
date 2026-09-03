package falseshare

import (
	"testing"
	"unsafe"
)

func TestCountTotal(t *testing.T) {
	if got := Count(4, 1000); got != 4000 {
		t.Errorf("Count = %d, want 4000", got)
	}
	if got := Count(1, 0); got != 0 {
		t.Errorf("Count = %d, want 0", got)
	}
	if got := Count(0, 10); got != 0 {
		t.Errorf("Count = %d, want 0", got)
	}
}

func TestCountUnderLoad(t *testing.T) {
	if got := Count(8, 100000); got != 800000 {
		t.Errorf("Count = %d, want 800000", got)
	}
}

func TestCountersDoNotShareALine(t *testing.T) {
	if got := unsafe.Sizeof(counter{}); got < cacheLine {
		t.Errorf("sizeof(counter) = %d, want at least %d: neighbouring counters share a cache line", got, cacheLine)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(4, 10000)
	}
}
