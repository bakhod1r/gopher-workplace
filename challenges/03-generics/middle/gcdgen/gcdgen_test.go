package gcdgen

import "testing"

func TestGCD(t *testing.T) {
	if got := GCD(12, 18); got != 6 {
		t.Errorf("GCD(12, 18) = %v, want 6", got)
	}
	if got := GCD(18, 12); got != 6 {
		t.Errorf("GCD(18, 12) = %v, want 6", got)
	}
	if got := GCD(7, 13); got != 1 {
		t.Errorf("GCD(7, 13) = %v, want 1", got)
	}
	if got := GCD(5, 0); got != 5 {
		t.Errorf("GCD(5, 0) = %v, want 5", got)
	}
	if got := GCD(0, 0); got != 0 {
		t.Errorf("GCD(0, 0) = %v, want 0", got)
	}
}

func TestGCDNegatives(t *testing.T) {
	if got := GCD(-12, 18); got != 6 {
		t.Errorf("GCD(-12, 18) = %v, want 6", got)
	}
	if got := GCD(-12, -18); got != 6 {
		t.Errorf("GCD(-12, -18) = %v, want 6", got)
	}
	if got := GCD(int64(-8), int64(4)); got != 4 {
		t.Errorf("GCD(-8, 4) = %v, want 4", got)
	}
}
