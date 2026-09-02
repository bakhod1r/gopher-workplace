package numberfull

import "testing"

func TestTotal(t *testing.T) {
	if got := Total([]int{1, -2}); got != -1 {
		t.Errorf("Total([]int{1, -2}) = %v, want -1", got)
	}
	if got := Total([]uint{1, 2}); got != 3 {
		t.Errorf("Total([]uint{1, 2}) = %v, want 3", got)
	}
	if got := Total([]uint64{5}); got != 5 {
		t.Errorf("Total([]uint64{5}) = %v, want 5", got)
	}
	if got := Total([]float64{0.5, 0.5}); got != 1 {
		t.Errorf("Total([]float64{...}) = %v, want 1", got)
	}
	if got := Total([]float32{1.5}); got != 1.5 {
		t.Errorf("Total([]float32{1.5}) = %v, want 1.5", got)
	}
	if got := Total([]int{}); got != 0 {
		t.Errorf("Total([]int{}) = %v, want 0", got)
	}
}
