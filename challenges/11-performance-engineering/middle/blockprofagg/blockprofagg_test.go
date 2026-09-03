package blockprofagg

import (
	"math"
	"reflect"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestScale(t *testing.T) {
	cases := []struct {
		wait int64
		rate int
		want int64
	}{
		{100, 1_000_000, 100_000_000},
		{100, 1, 100},
		{100, 0, 100},
		{100, -5, 100},
		{0, 1000, 0},
	}
	for _, c := range cases {
		if got := Scale(c.wait, c.rate); got != c.want {
			t.Errorf("Scale(%d, %d) = %d, want %d", c.wait, c.rate, got, c.want)
		}
	}
}

func TestTotals(t *testing.T) {
	got := Totals([]Event{{"a", 100}, {"b", 20}, {"a", 50}}, 1)
	want := map[string]int64{"a": 150, "b": 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Totals = %v, want %v", got, want)
	}
}

func TestTotalsAppliesTheRate(t *testing.T) {
	got := Totals([]Event{{"a", 100}}, 1000)
	if got["a"] != 100_000 {
		t.Errorf("Totals[a] = %d, want 100000 — the sampled wait stands for rate times as much", got["a"])
	}
}

func TestTotalsDropsJunk(t *testing.T) {
	got := Totals([]Event{{"a", 0}, {"b", -5}}, 1)
	if got == nil || len(got) != 0 {
		t.Errorf("Totals = %v, want empty non-nil map", got)
	}
}

func TestFractionBlocked(t *testing.T) {
	if got := FractionBlocked([]Event{{"a", 50}}, 1, 100); !near(got, 0.5) {
		t.Errorf("FractionBlocked = %v, want 0.5", got)
	}
	// Two goroutines blocked for the whole window: 2.0, not 1.0.
	if got := FractionBlocked([]Event{{"a", 100}, {"b", 100}}, 1, 100); !near(got, 2) {
		t.Errorf("FractionBlocked = %v, want 2 — goroutines block in parallel", got)
	}
	if got := FractionBlocked([]Event{{"a", 50}}, 1, 0); got != 0 {
		t.Errorf("FractionBlocked with a zero window = %v, want 0", got)
	}
	if got := FractionBlocked(nil, 1, 100); got != 0 {
		t.Errorf("FractionBlocked(nil) = %v, want 0", got)
	}
}
