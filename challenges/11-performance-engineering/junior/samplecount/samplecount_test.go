package samplecount

import "testing"

func TestTotals(t *testing.T) {
	count, nanos := Totals([]Sample{{"a", 3, 10}, {"b", 2, 10}})
	if count != 5 || nanos != 50 {
		t.Errorf("Totals = %d, %d, want 5, 50", count, nanos)
	}
}

func TestTotalsMixedPeriods(t *testing.T) {
	count, nanos := Totals([]Sample{{"a", 2, 10_000_000}, {"b", 1, 5_000_000}})
	if count != 3 || nanos != 25_000_000 {
		t.Errorf("Totals = %d, %d, want 3, 25000000", count, nanos)
	}
}

func TestTotalsSkipsJunk(t *testing.T) {
	count, nanos := Totals([]Sample{{"a", 0, 10}, {"b", 3, 0}, {"c", -1, 10}, {"d", 2, -5}})
	if count != 0 || nanos != 0 {
		t.Errorf("Totals = %d, %d, want 0, 0", count, nanos)
	}
}

func TestTotalsEmpty(t *testing.T) {
	count, nanos := Totals(nil)
	if count != 0 || nanos != 0 {
		t.Errorf("Totals(nil) = %d, %d, want 0, 0", count, nanos)
	}
}
