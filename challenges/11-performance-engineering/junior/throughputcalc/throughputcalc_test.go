package throughputcalc

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-6 }

func TestOpsPerSec(t *testing.T) {
	cases := []struct{ ns, want float64 }{
		{1000, 1e6},
		{1, 1e9},
		{1e9, 1},
		{0, 0},
		{-1, 0},
	}
	for _, c := range cases {
		if got := OpsPerSec(c.ns); !near(got, c.want) {
			t.Errorf("OpsPerSec(%v) = %v, want %v", c.ns, got, c.want)
		}
	}
}

func TestCapacity(t *testing.T) {
	cases := []struct {
		ns    float64
		cores int
		want  int64
	}{
		{1000, 8, 8_000_000},
		{1000, 1, 1_000_000},
		{1e9, 4, 4},
		{1000, 0, 0},
		{1000, -2, 0},
		{0, 8, 0},
	}
	for _, c := range cases {
		if got := Capacity(c.ns, c.cores); got != c.want {
			t.Errorf("Capacity(%v, %d) = %d, want %d", c.ns, c.cores, got, c.want)
		}
	}
}

func TestCapacityTruncates(t *testing.T) {
	// 1.5 ops/sec per core, 3 cores -> 4.5 -> 4
	if got := Capacity(2e9/3, 3); got != 4 {
		t.Errorf("Capacity = %d, want 4", got)
	}
}
