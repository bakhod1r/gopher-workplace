package heapgrowthrate

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-6 }

func TestNextTarget(t *testing.T) {
	cases := []struct {
		live int64
		gogc int
		want int64
		ok   bool
	}{
		{4 << 20, 100, 8 << 20, true},
		{4 << 20, 50, 6 << 20, true},
		{4 << 20, 200, 12 << 20, true},
		{0, 100, 0, true},
		{4 << 20, 0, 0, false},
		{4 << 20, -1, 0, false},
	}
	for _, c := range cases {
		got, ok := NextTarget(c.live, c.gogc)
		if got != c.want || ok != c.ok {
			t.Errorf("NextTarget(%d, %d) = %d, %v, want %d, %v", c.live, c.gogc, got, ok, c.want, c.ok)
		}
	}
}

func TestGrowthPerSec(t *testing.T) {
	got, ok := GrowthPerSec([]Sample{{0, 100}, {1_000_000_000, 300}})
	if !ok || !near(got, 200) {
		t.Errorf("GrowthPerSec = %v, %v, want 200, true", got, ok)
	}
	got, ok = GrowthPerSec([]Sample{{0, 300}, {2_000_000_000, 100}})
	if !ok || !near(got, -100) {
		t.Errorf("GrowthPerSec = %v, %v, want -100, true — a shrinking heap has a negative slope", got, ok)
	}
}

func TestGrowthPerSecUsesFirstAndLast(t *testing.T) {
	// The sawtooth in between must not change the endpoints' slope.
	got, ok := GrowthPerSec([]Sample{
		{0, 100},
		{500_000_000, 900},
		{1_000_000_000, 300},
	})
	if !ok || !near(got, 200) {
		t.Errorf("GrowthPerSec = %v, %v, want 200, true", got, ok)
	}
}

func TestGrowthPerSecGuards(t *testing.T) {
	if _, ok := GrowthPerSec(nil); ok {
		t.Error("GrowthPerSec(nil) reported ok")
	}
	if _, ok := GrowthPerSec([]Sample{{0, 100}}); ok {
		t.Error("GrowthPerSec with one sample reported ok")
	}
	if _, ok := GrowthPerSec([]Sample{{100, 1}, {100, 2}}); ok {
		t.Error("GrowthPerSec with a zero time span reported ok")
	}
	if _, ok := GrowthPerSec([]Sample{{200, 1}, {100, 2}}); ok {
		t.Error("GrowthPerSec with time going backwards reported ok")
	}
}

func TestDoubling(t *testing.T) {
	got, ok := Doubling([]Sample{{0, 100}, {1_000_000_000, 200}})
	if !ok || !near(got, 2) {
		t.Errorf("Doubling = %v, %v, want 2, true", got, ok)
	}
	if _, ok := Doubling([]Sample{{0, 300}, {1_000_000_000, 100}}); ok {
		t.Error("a shrinking heap reported a doubling time")
	}
	if _, ok := Doubling([]Sample{{0, 100}, {1_000_000_000, 100}}); ok {
		t.Error("a flat heap reported a doubling time")
	}
	if _, ok := Doubling(nil); ok {
		t.Error("Doubling(nil) reported ok")
	}
}
