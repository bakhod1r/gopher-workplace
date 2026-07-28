package bumpfield

import "testing"

func TestGrow(t *testing.T) {
	c := Cart{Count: 2}
	Grow(&c)
	Grow(&c)
	if c.Count != 4 {
		t.Errorf("Count=%d want 4", c.Count)
	}
}
