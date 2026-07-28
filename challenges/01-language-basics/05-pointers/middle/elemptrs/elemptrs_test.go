package elemptrs

import "testing"

func TestPointers(t *testing.T) {
	xs := []int{10, 20, 30}
	ps := Pointers(xs)
	*ps[1] = 99
	if xs[1] != 99 {
		t.Errorf("xs[1]=%d want 99 (pointers must alias xs)", xs[1])
	}
	if len(ps) != 3 {
		t.Errorf("len=%d want 3", len(ps))
	}
}
