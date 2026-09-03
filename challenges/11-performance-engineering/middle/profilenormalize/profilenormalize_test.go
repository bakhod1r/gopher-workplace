package profilenormalize

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-9 }

func TestRate(t *testing.T) {
	got := Rate(map[string]int64{"a": 60, "b": 30}, 30)
	if len(got) != 2 || !near(got["a"], 2) || !near(got["b"], 1) {
		t.Errorf("Rate = %v, want {a:2 b:1}", got)
	}
}

func TestRateMakesDifferentDurationsComparable(t *testing.T) {
	short := Rate(map[string]int64{"a": 10}, 5)
	long := Rate(map[string]int64{"a": 40}, 20)
	if !near(short["a"], long["a"]) {
		t.Errorf("rates differ: %v vs %v — same workload, different durations", short["a"], long["a"])
	}
}

func TestRateGuards(t *testing.T) {
	for _, s := range []float64{0, -1} {
		got := Rate(map[string]int64{"a": 60}, s)
		if got == nil || len(got) != 0 {
			t.Errorf("Rate(_, %v) = %v, want empty non-nil map", s, got)
		}
	}
}

func TestRateDoesNotModifyInput(t *testing.T) {
	in := map[string]int64{"a": 60}
	Rate(in, 30)
	if in["a"] != 60 {
		t.Errorf("input changed: %v", in)
	}
}

func TestFractions(t *testing.T) {
	got := Fractions(map[string]int64{"a": 3, "b": 1})
	if !near(got["a"], 0.75) || !near(got["b"], 0.25) {
		t.Errorf("Fractions = %v, want {a:0.75 b:0.25}", got)
	}
	sum := 0.0
	for _, v := range got {
		sum += v
	}
	if !near(sum, 1) {
		t.Errorf("fractions sum to %v, want 1", sum)
	}
}

func TestFractionsDropsNonPositive(t *testing.T) {
	got := Fractions(map[string]int64{"a": 3, "b": 0, "c": -5})
	if len(got) != 1 || !near(got["a"], 1) {
		t.Errorf("Fractions = %v, want {a:1}", got)
	}
}

func TestFractionsEmpty(t *testing.T) {
	for _, in := range []map[string]int64{nil, {"a": 0}} {
		got := Fractions(in)
		if got == nil || len(got) != 0 {
			t.Errorf("Fractions(%v) = %v, want empty non-nil map", in, got)
		}
	}
}
