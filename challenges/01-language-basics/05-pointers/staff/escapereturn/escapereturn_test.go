package escapereturn

import "testing"

func TestMaxPtr(t *testing.T) {
	xs := []int{3, 9, 4}
	p := MaxPtr(xs)
	*p = 0
	if xs[1] != 0 {
		t.Errorf("xs[1]=%d want 0 (pointer must alias the slice element)", xs[1])
	}
}
