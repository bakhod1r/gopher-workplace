package counter

import "testing"

func TestCounter(t *testing.T) {
	c := MakeCounter()
	if c() != 1 || c() != 2 || c() != 3 {
		t.Errorf("counter did not increment 1,2,3")
	}
	d := MakeCounter()
	if d() != 1 {
		t.Errorf("second counter not independent")
	}
}
