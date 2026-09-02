package productgen

import "testing"

func TestProduct(t *testing.T) {
	if got := Product([]int{2, 3}); got != 6 {
		t.Errorf("Product([]int{2, 3}) = %v, want 6", got)
	}
	if got := Product([]float64{2, 0.5}); got != 1 {
		t.Errorf("Product([]float64{2, 0.5}) = %v, want 1", got)
	}
	if got := Product([]int{}); got != 1 {
		t.Errorf("Product([]int{}) = %v, want 1", got)
	}
	if got := Product([]int{5, 0}); got != 0 {
		t.Errorf("Product([]int{5, 0}) = %v, want 0", got)
	}
}
