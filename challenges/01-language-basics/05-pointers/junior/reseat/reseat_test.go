package reseat

import "testing"

func TestReseat(t *testing.T) {
	a, b := 1, 2
	p := &a
	Reseat(&p, &b)
	if p != &b {
		t.Errorf("p should now point to b")
	}
}
