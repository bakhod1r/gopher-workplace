package accgen

import "testing"

func TestAcc(t *testing.T) {
	var a Acc[int]
	if got := a.Mean(); got != 0 {
		t.Errorf("Mean() before Add = %v, want 0", got)
	}
	a.Add(1)
	a.Add(2)
	if got := a.Sum(); got != 3 {
		t.Errorf("Sum() = %v, want 3", got)
	}
	if got := a.Mean(); got != 1.5 {
		t.Errorf("Mean() = %v, want 1.5 (do not truncate)", got)
	}
}

func TestAccFloats(t *testing.T) {
	var a Acc[float64]
	a.Add(0.5)
	a.Add(1.5)
	if got := a.Sum(); got != 2 {
		t.Errorf("Sum() = %v, want 2", got)
	}
	if got := a.Mean(); got != 1 {
		t.Errorf("Mean() = %v, want 1", got)
	}
}
