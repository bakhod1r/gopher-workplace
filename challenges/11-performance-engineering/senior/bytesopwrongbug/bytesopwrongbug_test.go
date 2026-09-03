package bytesopwrongbug

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestThroughputUsesDecimalMegabytes(t *testing.T) {
	if got := ThroughputMBs(1_000_000, 1_000_000_000); !near(got, 1) {
		t.Errorf("ThroughputMBs = %v, want exactly 1 — the tool uses 1e6 bytes per MB", got)
	}
}

func TestThroughputScales(t *testing.T) {
	cases := []struct {
		bytes, ns int64
		want      float64
	}{
		{2_000_000, 1_000_000_000, 2},
		{1_000_000, 500_000_000, 2},
		{500, 1_000, 500},
		{0, 1_000_000_000, 0},
	}
	for _, c := range cases {
		if got := ThroughputMBs(c.bytes, c.ns); !near(got, c.want) {
			t.Errorf("ThroughputMBs(%d, %d) = %v, want %v", c.bytes, c.ns, got, c.want)
		}
	}
}

func TestThroughputGuards(t *testing.T) {
	for _, ns := range []int64{0, -1} {
		if got := ThroughputMBs(1_000_000, ns); got != 0 {
			t.Errorf("ThroughputMBs(_, %d) = %v, want 0", ns, got)
		}
	}
}

func TestPerOpBytes(t *testing.T) {
	if got := PerOpBytes(2048, 4); got != 512 {
		t.Errorf("PerOpBytes = %d, want 512", got)
	}
	if got := PerOpBytes(2048, 0); got != 0 {
		t.Errorf("PerOpBytes = %d, want 0", got)
	}
}
