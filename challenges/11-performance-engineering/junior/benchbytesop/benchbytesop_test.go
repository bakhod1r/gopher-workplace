package benchbytesop

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestThroughput(t *testing.T) {
	cases := []struct {
		bytes, ns int64
		want      float64
	}{
		{1_000_000, 1_000_000_000, 1},
		{2_000_000, 1_000_000_000, 2},
		{1_000_000, 500_000_000, 2},
		{0, 1_000_000_000, 0},
		{500, 1_000, 500},
	}
	for _, c := range cases {
		if got := ThroughputMBs(c.bytes, c.ns); !near(got, c.want) {
			t.Errorf("ThroughputMBs(%d, %d) = %v, want %v", c.bytes, c.ns, got, c.want)
		}
	}
}

func TestThroughputNonPositiveElapsed(t *testing.T) {
	for _, ns := range []int64{0, -1} {
		if got := ThroughputMBs(1_000_000, ns); got != 0 {
			t.Errorf("ThroughputMBs(_, %d) = %v, want 0", ns, got)
		}
	}
}
