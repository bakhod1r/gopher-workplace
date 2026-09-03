package gcpausesum

import (
	"math"
	"testing"
)

func near(got, want float64) bool { return math.Abs(got-want) < 1e-12 }

func TestTotal(t *testing.T) {
	if got := Total([]int64{100, 200}); got != 300 {
		t.Errorf("Total = %d, want 300", got)
	}
	if got := Total([]int64{100, -50, 200}); got != 300 {
		t.Errorf("Total = %d, want 300 — negative pauses are ignored", got)
	}
	if got := Total(nil); got != 0 {
		t.Errorf("Total(nil) = %d, want 0", got)
	}
}

func TestWorst(t *testing.T) {
	v, i := Worst([]int64{100, 500, 200})
	if v != 500 || i != 1 {
		t.Errorf("Worst = %d, %d, want 500, 1", v, i)
	}
	v, i = Worst([]int64{7, 7, 7})
	if v != 7 || i != 0 {
		t.Errorf("Worst = %d, %d, want 7, 0", v, i)
	}
	v, i = Worst(nil)
	if v != 0 || i != -1 {
		t.Errorf("Worst(nil) = %d, %d, want 0, -1", v, i)
	}
	v, i = Worst([]int64{-5})
	if v != 0 || i != -1 {
		t.Errorf("Worst([-5]) = %d, %d, want 0, -1", v, i)
	}
}

func TestFractionOf(t *testing.T) {
	if got := FractionOf([]int64{5_000_000}, 1_000_000_000); !near(got, 0.005) {
		t.Errorf("FractionOf = %v, want 0.005", got)
	}
	if got := FractionOf([]int64{5_000_000}, 10_000_000); !near(got, 0.5) {
		t.Errorf("FractionOf = %v, want 0.5 — the same pause in a shorter window", got)
	}
	if got := FractionOf([]int64{1}, 0); got != 0 {
		t.Errorf("FractionOf with a zero window = %v, want 0", got)
	}
}

func TestWithinBudget(t *testing.T) {
	pauses := []int64{1_000_000, 2_000_000}
	if !WithinBudget(pauses, 1_000_000_000, 0.01, 5_000_000) {
		t.Error("WithinBudget = false, want true")
	}
	if WithinBudget(pauses, 1_000_000_000, 0.001, 5_000_000) {
		t.Error("WithinBudget = true, want false — the total fraction is over budget")
	}
	if WithinBudget(pauses, 1_000_000_000, 0.01, 1_500_000) {
		t.Error("WithinBudget = true, want false — one pause exceeds the per-pause limit")
	}
}

func TestWithinBudgetBoundariesAreInclusive(t *testing.T) {
	if !WithinBudget([]int64{5_000_000}, 1_000_000_000, 0.005, 5_000_000) {
		t.Error("exactly at both limits reported over budget; the limits are inclusive")
	}
}

func TestWithinBudgetNoPauses(t *testing.T) {
	if !WithinBudget(nil, 1_000_000_000, 0.0, 0) {
		t.Error("no pauses reported over budget")
	}
}
