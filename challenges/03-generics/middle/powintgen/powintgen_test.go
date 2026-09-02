package powintgen

import "testing"

func TestPow(t *testing.T) {
	if got := Pow(2, 10); got != 1024 {
		t.Errorf("Pow(2, 10) = %v, want 1024", got)
	}
	if got := Pow(3, 3); got != 27 {
		t.Errorf("Pow(3, 3) = %v, want 27", got)
	}
	if got := Pow(5, 1); got != 5 {
		t.Errorf("Pow(5, 1) = %v, want 5", got)
	}
	if got := Pow(7, 0); got != 1 {
		t.Errorf("Pow(7, 0) = %v, want 1", got)
	}
	if got := Pow(int64(2), 5); got != 32 {
		t.Errorf("Pow(int64(2), 5) = %v, want 32", got)
	}
}

func TestPowNegativeExponent(t *testing.T) {
	if got := Pow(2, -1); got != 0 {
		t.Errorf("Pow(2, -1) = %v, want 0", got)
	}
}

func TestPowNegativeBase(t *testing.T) {
	if got := Pow(-2, 3); got != -8 {
		t.Errorf("Pow(-2, 3) = %v, want -8", got)
	}
	if got := Pow(-2, 2); got != 4 {
		t.Errorf("Pow(-2, 2) = %v, want 4", got)
	}
}
