package roundgen

import "testing"

func TestRoundHalfUp(t *testing.T) {
	if got := RoundHalfUp(2.5); got != 3 {
		t.Errorf("RoundHalfUp(2.5) = %v, want 3", got)
	}
	if got := RoundHalfUp(2.4); got != 2 {
		t.Errorf("RoundHalfUp(2.4) = %v, want 2", got)
	}
	if got := RoundHalfUp(-2.5); got != -3 {
		t.Errorf("RoundHalfUp(-2.5) = %v, want -3 (away from zero)", got)
	}
	if got := RoundHalfUp(-2.4); got != -2 {
		t.Errorf("RoundHalfUp(-2.4) = %v, want -2", got)
	}
	if got := RoundHalfUp(3.5); got != 4 {
		t.Errorf("RoundHalfUp(3.5) = %v, want 4 (not banker's rounding)", got)
	}
	if got := RoundHalfUp(float32(1.5)); got != 2 {
		t.Errorf("RoundHalfUp(float32) = %v, want 2", got)
	}
}
