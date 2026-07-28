package samecart

import "testing"

func TestSame(t *testing.T) {
	c := &Cart{Count: 1}
	d := &Cart{Count: 1}
	if !Same(c, c) {
		t.Errorf("c,c same")
	}
	if Same(c, d) {
		t.Errorf("c,d distinct instances")
	}
}
