package eqconfuse

import "testing"

func TestSame(t *testing.T) {
	x := 5
	y := 5
	if Same(&x, &y) {
		t.Errorf("distinct vars with equal values are not the same pointer")
	}
	if !Same(&x, &x) {
		t.Errorf("&x,&x should be the same")
	}
}
