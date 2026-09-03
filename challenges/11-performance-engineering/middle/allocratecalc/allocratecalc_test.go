package allocratecalc

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-6 }

func TestDelta(t *testing.T) {
	a := Stats{NS: 0, TotalAlloc: 100, Mallocs: 10, Frees: 4}
	b := Stats{NS: 1_000_000_000, TotalAlloc: 500, Mallocs: 30, Frees: 20}
	bytes, mallocs, frees, ok := Delta(a, b)
	if !ok || bytes != 400 || mallocs != 20 || frees != 16 {
		t.Errorf("Delta = %d, %d, %d, %v, want 400, 20, 16, true", bytes, mallocs, frees, ok)
	}
}

func TestDeltaRejectsCountersGoingBackwards(t *testing.T) {
	a := Stats{NS: 0, TotalAlloc: 500, Mallocs: 30, Frees: 20}
	b := Stats{NS: 1_000_000_000, TotalAlloc: 100, Mallocs: 10, Frees: 4}
	if _, _, _, ok := Delta(a, b); ok {
		t.Error("Delta accepted a decreasing cumulative counter; that means the process restarted")
	}
}

func TestDeltaRejectsBadTimestamps(t *testing.T) {
	a := Stats{NS: 1_000_000_000, TotalAlloc: 100}
	b := Stats{NS: 0, TotalAlloc: 200}
	if _, _, _, ok := Delta(a, b); ok {
		t.Error("Delta accepted time going backwards")
	}
	if _, _, _, ok := Delta(a, Stats{NS: a.NS, TotalAlloc: 200}); ok {
		t.Error("Delta accepted two snapshots at the same instant")
	}
}

func TestBytesPerSec(t *testing.T) {
	a := Stats{NS: 0, TotalAlloc: 0}
	b := Stats{NS: 2_000_000_000, TotalAlloc: 2_000_000}
	got, ok := BytesPerSec(a, b)
	if !ok || !near(got, 1_000_000) {
		t.Errorf("BytesPerSec = %v, %v, want 1000000, true", got, ok)
	}
	if _, ok := BytesPerSec(b, a); ok {
		t.Error("BytesPerSec accepted reversed snapshots")
	}
}

func TestLiveObjects(t *testing.T) {
	if got, ok := LiveObjects(Stats{Mallocs: 10, Frees: 4}); !ok || got != 6 {
		t.Errorf("LiveObjects = %d, %v, want 6, true", got, ok)
	}
	if got, ok := LiveObjects(Stats{Mallocs: 5, Frees: 5}); !ok || got != 0 {
		t.Errorf("LiveObjects = %d, %v, want 0, true", got, ok)
	}
	if _, ok := LiveObjects(Stats{Mallocs: 4, Frees: 10}); ok {
		t.Error("LiveObjects accepted more frees than mallocs")
	}
}

func TestUnsignedSubtractionDoesNotWrap(t *testing.T) {
	// The classic bug: uint64(4) - uint64(10) is about 1.8e19, not -6.
	if got, ok := LiveObjects(Stats{Mallocs: 4, Frees: 10}); ok || got != 0 {
		t.Errorf("LiveObjects = %d, %v, want 0, false — an unsigned subtraction must be guarded, not wrapped", got, ok)
	}
}
