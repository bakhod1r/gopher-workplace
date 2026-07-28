package arrayvalue

import "testing"

func TestDouble(t *testing.T) {
	a := [3]int{1, 2, 3}
	Double(&a)
	if a != [3]int{2, 4, 6} {
		t.Errorf("a=%v; want [2 4 6]", a)
	}
}
