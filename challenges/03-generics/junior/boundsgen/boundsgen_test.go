package boundsgen

import "testing"

func TestBounds(t *testing.T) {
	lo, hi, ok := Bounds([]int{3, 1, 2})
	if lo != 1 || hi != 3 || !ok {
		t.Errorf("Bounds([]int{3, 1, 2}) = %v, %v, %v, want 1, 3, true", lo, hi, ok)
	}
	lo, hi, ok = Bounds([]int{5})
	if lo != 5 || hi != 5 || !ok {
		t.Errorf("Bounds([]int{5}) = %v, %v, %v, want 5, 5, true", lo, hi, ok)
	}
	slo, shi, ok := Bounds([]string{"b", "a", "c"})
	if slo != "a" || shi != "c" || !ok {
		t.Errorf("Bounds([]string{...}) = %q, %q, %v, want a, c, true", slo, shi, ok)
	}
	lo, hi, ok = Bounds([]int{})
	if lo != 0 || hi != 0 || ok {
		t.Errorf("Bounds([]int{}) = %v, %v, %v, want 0, 0, false", lo, hi, ok)
	}
}
