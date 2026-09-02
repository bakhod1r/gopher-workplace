package absgen

import "testing"

func TestAbs(t *testing.T) {
	if got := Abs(-3); got != 3 {
		t.Errorf("Abs(-3) = %v, want 3", got)
	}
	if got := Abs(3); got != 3 {
		t.Errorf("Abs(3) = %v, want 3", got)
	}
	if got := Abs(0); got != 0 {
		t.Errorf("Abs(0) = %v, want 0", got)
	}
	if got := Abs(-2.5); got != 2.5 {
		t.Errorf("Abs(-2.5) = %v, want 2.5", got)
	}
	if got := Abs(int64(-7)); got != 7 {
		t.Errorf("Abs(int64(-7)) = %v, want 7", got)
	}
}
