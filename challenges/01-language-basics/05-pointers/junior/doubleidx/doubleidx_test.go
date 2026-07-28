package doubleidx

import "testing"

func TestDouble(t *testing.T) {
	a := [4]int{1, 2, 3, 4}
	Double(&a, 2)
	if a[2] != 6 {
		t.Errorf("a[2]=%d want 6", a[2])
	}
}
