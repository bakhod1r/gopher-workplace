package trackergen

import "testing"

func TestTracker(t *testing.T) {
	var tr Tracker[int]
	if lo, hi, ok := tr.Bounds(); lo != 0 || hi != 0 || ok {
		t.Errorf("Bounds() before Add = %v, %v, %v, want 0, 0, false", lo, hi, ok)
	}
	tr.Add(3)
	if lo, hi, ok := tr.Bounds(); lo != 3 || hi != 3 || !ok {
		t.Errorf("Bounds() = %v, %v, %v, want 3, 3, true", lo, hi, ok)
	}
	tr.Add(1)
	tr.Add(5)
	if lo, hi, ok := tr.Bounds(); lo != 1 || hi != 5 || !ok {
		t.Errorf("Bounds() = %v, %v, %v, want 1, 5, true", lo, hi, ok)
	}
}

func TestTrackerAllPositive(t *testing.T) {
	var tr Tracker[int]
	tr.Add(4)
	tr.Add(7)
	if lo, _, _ := tr.Bounds(); lo != 4 {
		t.Errorf("low bound = %v, want 4 (do not seed from the zero value)", lo)
	}
}

func TestTrackerStrings(t *testing.T) {
	var tr Tracker[string]
	tr.Add("b")
	tr.Add("a")
	if lo, hi, ok := tr.Bounds(); lo != "a" || hi != "b" || !ok {
		t.Errorf("Bounds() = %q, %q, %v, want a, b, true", lo, hi, ok)
	}
}
