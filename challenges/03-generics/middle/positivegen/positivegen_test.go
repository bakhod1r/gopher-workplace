package positivegen

import "testing"

func TestAllPositive(t *testing.T) {
	if !AllPositive([]int{1, 2}) {
		t.Error("AllPositive([1 2]) = false, want true")
	}
	if AllPositive([]int{1, 0}) {
		t.Error("AllPositive([1 0]) = true, want false (zero is not positive)")
	}
	if AllPositive([]int{-1}) {
		t.Error("AllPositive([-1]) = true, want false")
	}
	if !AllPositive([]int{}) {
		t.Error("AllPositive([]) = false, want true")
	}
}

func TestFirstNonPositive(t *testing.T) {
	if got := FirstNonPositive([]int{1, -1}); got != 1 {
		t.Errorf("FirstNonPositive = %d, want 1", got)
	}
	if got := FirstNonPositive([]int{0, 1}); got != 0 {
		t.Errorf("FirstNonPositive = %d, want 0", got)
	}
	if got := FirstNonPositive([]int{1, 2}); got != -1 {
		t.Errorf("FirstNonPositive = %d, want -1", got)
	}
	if got := FirstNonPositive([]int64{}); got != -1 {
		t.Errorf("FirstNonPositive(empty) = %d, want -1", got)
	}
}
