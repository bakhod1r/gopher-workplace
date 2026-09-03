package stdequalfuncfieldbug

import "testing"

func TestSameLinesIgnoresNote(t *testing.T) {
	a := []Line{{"x", 1, "picked"}}
	b := []Line{{"x", 1, ""}}
	if !SameLines(a, b) {
		t.Errorf("SameLines = false, want true (Note must be ignored)")
	}
}

func TestSameLinesSeesQtyAndSKU(t *testing.T) {
	if SameLines([]Line{{"x", 1, "a"}}, []Line{{"x", 2, "a"}}) {
		t.Errorf("SameLines = true, want false (Qty differs)")
	}
	if SameLines([]Line{{"x", 1, "a"}}, []Line{{"y", 1, "a"}}) {
		t.Errorf("SameLines = true, want false (SKU differs)")
	}
}

func TestSameLinesLength(t *testing.T) {
	if SameLines([]Line{{"x", 1, ""}}, nil) {
		t.Errorf("SameLines = true, want false (length differs)")
	}
	if !SameLines(nil, []Line{}) {
		t.Errorf("SameLines = false, want true (both empty)")
	}
}
