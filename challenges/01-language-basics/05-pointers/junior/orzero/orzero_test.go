package orzero

import "testing"

func TestDerefOrZero(t *testing.T) {
	if DerefOrZero(nil) != 0 {
		t.Errorf("nil should be 0")
	}
	x := 8
	if DerefOrZero(&x) != 8 {
		t.Errorf("want 8")
	}
}
