package valueor

import "testing"

func TestValueOr(t *testing.T) {
	if ValueOr(nil, 5) != 5 {
		t.Errorf("nil should give default")
	}
	x := 9
	if ValueOr(&x, 5) != 9 {
		t.Errorf("want 9")
	}
}
