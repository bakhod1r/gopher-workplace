package reportallocsscopebug

import "testing"

func TestPerOpDividesByMeasuredIterations(t *testing.T) {
	b, a := PerOp(Run{Warmup: 90, Measured: 10, Bytes: 800, Allocs: 20})
	if b != 80 || a != 2 {
		t.Errorf("PerOp = %d, %d, want 80, 2 — warmup iterations are not covered by the counters", b, a)
	}
}

func TestPerOpWithoutWarmup(t *testing.T) {
	b, a := PerOp(Run{Warmup: 0, Measured: 4, Bytes: 2048, Allocs: 8})
	if b != 512 || a != 2 {
		t.Errorf("PerOp = %d, %d, want 512, 2", b, a)
	}
}

func TestPerOpIsIndependentOfWarmupCount(t *testing.T) {
	base, _ := PerOp(Run{Warmup: 0, Measured: 10, Bytes: 1000, Allocs: 10})
	for _, w := range []int64{1, 100, 1_000_000} {
		got, _ := PerOp(Run{Warmup: w, Measured: 10, Bytes: 1000, Allocs: 10})
		if got != base {
			t.Fatalf("warmup %d changed B/op to %d, want %d", w, got, base)
		}
	}
	if base != 100 {
		t.Errorf("B/op = %d, want 100", base)
	}
}

func TestPerOpNoMeasuredIterations(t *testing.T) {
	b, a := PerOp(Run{Warmup: 100, Measured: 0, Bytes: 999, Allocs: 99})
	if b != 0 || a != 0 {
		t.Errorf("PerOp = %d, %d, want 0, 0", b, a)
	}
}
