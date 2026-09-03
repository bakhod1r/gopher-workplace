package gcbudget

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestCPUFraction(t *testing.T) {
	if got := CPUFraction(100_000_000, 1_000_000_000, 1); !near(got, 0.1) {
		t.Errorf("CPUFraction = %v, want 0.1", got)
	}
	// The same GC time spread over 8 cores is a smaller share of total CPU.
	if got := CPUFraction(100_000_000, 1_000_000_000, 8); !near(got, 0.0125) {
		t.Errorf("CPUFraction = %v, want 0.0125", got)
	}
	if got := CPUFraction(100, 0, 4); got != 0 {
		t.Errorf("CPUFraction with a zero window = %v, want 0", got)
	}
	if got := CPUFraction(100, 1000, 0); got != 0 {
		t.Errorf("CPUFraction with zero cores = %v, want 0", got)
	}
}

func TestTuneGOGC(t *testing.T) {
	got, ok := TuneGOGC(100, 0.20, 0.10, 1000)
	if !ok || got != 200 {
		t.Errorf("TuneGOGC = %d, %v, want 200, true", got, ok)
	}
	got, ok = TuneGOGC(100, 0.30, 0.10, 1000)
	if !ok || got != 300 {
		t.Errorf("TuneGOGC = %d, %v, want 300, true", got, ok)
	}
}

func TestTuneGOGCRespectsTheCeiling(t *testing.T) {
	got, ok := TuneGOGC(100, 0.50, 0.01, 400)
	if !ok || got != 400 {
		t.Errorf("TuneGOGC = %d, %v, want the ceiling 400, true", got, ok)
	}
}

func TestTuneGOGCWhenTheBudgetIsAlreadyMet(t *testing.T) {
	if _, ok := TuneGOGC(100, 0.05, 0.10, 1000); ok {
		t.Error("TuneGOGC suggested a change when the budget was already met")
	}
	if _, ok := TuneGOGC(100, 0.10, 0.10, 1000); ok {
		t.Error("TuneGOGC suggested a change when the fraction equals the target")
	}
}

func TestTuneGOGCGuards(t *testing.T) {
	if _, ok := TuneGOGC(0, 0.2, 0.1, 1000); ok {
		t.Error("TuneGOGC accepted a non-positive current GOGC")
	}
	if _, ok := TuneGOGC(100, 0.2, 0, 1000); ok {
		t.Error("TuneGOGC accepted a non-positive target")
	}
}

func TestMemoryCost(t *testing.T) {
	if got, ok := MemoryCost(100); !ok || !near(got, 2) {
		t.Errorf("MemoryCost(100) = %v, %v, want 2, true", got, ok)
	}
	if got, ok := MemoryCost(200); !ok || !near(got, 3) {
		t.Errorf("MemoryCost(200) = %v, %v, want 3, true", got, ok)
	}
	if got, ok := MemoryCost(50); !ok || !near(got, 1.5) {
		t.Errorf("MemoryCost(50) = %v, %v, want 1.5, true", got, ok)
	}
	if _, ok := MemoryCost(0); ok {
		t.Error("MemoryCost(0) reported ok")
	}
}

func TestTheTradeIsCPUForMemory(t *testing.T) {
	next, ok := TuneGOGC(100, 0.20, 0.10, 1000)
	if !ok {
		t.Fatal("TuneGOGC reported not ok")
	}
	before, _ := MemoryCost(100)
	after, _ := MemoryCost(next)
	if after <= before {
		t.Errorf("memory cost went from %v to %v; halving GC CPU must cost heap", before, after)
	}
}
