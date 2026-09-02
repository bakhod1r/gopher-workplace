package percentgen

import "testing"

func TestPercent(t *testing.T) {
	if got := Percent(1, 4); got != 25 {
		t.Errorf("Percent(1, 4) = %v, want 25 (convert before dividing)", got)
	}
	if got := Percent(3, 3); got != 100 {
		t.Errorf("Percent(3, 3) = %v, want 100", got)
	}
	if got := Percent(0.5, 2.0); got != 25 {
		t.Errorf("Percent(0.5, 2.0) = %v, want 25", got)
	}
	if got := Percent(1, 0); got != 0 {
		t.Errorf("Percent(1, 0) = %v, want 0", got)
	}
}
