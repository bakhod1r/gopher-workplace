package sameptr

import "testing"

func TestSame(t *testing.T) {
	x := 5
	y := 5
	if !Same(&x, &x) {
		t.Errorf("&x,&x should be same")
	}
	if Same(&x, &y) {
		t.Errorf("&x,&y are different addresses")
	}
}
