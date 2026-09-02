package unitmulgen

import "testing"

func TestTimes(t *testing.T) {
	if got := Times(Meters(2), 3); got != Meters(6) {
		t.Errorf("Times = %v, want 6", got)
	}
	if got := Times(Seconds(1.5), 2); got != Seconds(3) {
		t.Errorf("Times = %v, want 3", got)
	}
	if got := Times(Meters(2), 0); got != Meters(0) {
		t.Errorf("Times = %v, want 0", got)
	}
}

func TestSumUnits(t *testing.T) {
	if got := SumUnits([]Meters{1, 2}); got != Meters(3) {
		t.Errorf("SumUnits = %v, want 3", got)
	}
	if got := SumUnits([]Seconds{0.5, 0.5}); got != Seconds(1) {
		t.Errorf("SumUnits = %v, want 1", got)
	}
	if got := SumUnits([]Meters{}); got != Meters(0) {
		t.Errorf("SumUnits(empty) = %v, want 0", got)
	}
}

func TestUnitsKeepTheirType(t *testing.T) {
	var m Meters = Times(Meters(1), 4)
	if m != Meters(4) {
		t.Errorf("result = %v, want Meters(4)", m)
	}
}
