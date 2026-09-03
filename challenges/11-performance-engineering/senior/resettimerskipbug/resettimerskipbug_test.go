package resettimerskipbug

import "testing"

func TestResetDiscards(t *testing.T) {
	var tm Timer
	tm.Add(500)
	tm.Reset()
	if got := tm.Elapsed(); got != 0 {
		t.Errorf("Elapsed after Reset = %d, want 0", got)
	}
}

func TestResetTwice(t *testing.T) {
	var tm Timer
	tm.Add(100)
	tm.Reset()
	tm.Add(7)
	tm.Reset()
	if got := tm.Elapsed(); got != 0 {
		t.Errorf("Elapsed = %d, want 0", got)
	}
}

func TestAddAfterResetAccumulates(t *testing.T) {
	var tm Timer
	tm.Add(500)
	tm.Reset()
	tm.Add(3)
	tm.Add(4)
	if got := tm.Elapsed(); got != 7 {
		t.Errorf("Elapsed = %d, want 7", got)
	}
}

func TestBenchmarkExcludesSetup(t *testing.T) {
	if got := Benchmark(500, 7, 3); got != 21 {
		t.Errorf("Benchmark(500, 7, 3) = %d, want 21", got)
	}
	if got := Benchmark(1_000_000, 7, 1); got != 7 {
		t.Errorf("Benchmark = %d, want 7", got)
	}
}

func TestAddIgnoresNegative(t *testing.T) {
	var tm Timer
	tm.Add(-5)
	tm.Add(10)
	if got := tm.Elapsed(); got != 10 {
		t.Errorf("Elapsed = %d, want 10", got)
	}
}
