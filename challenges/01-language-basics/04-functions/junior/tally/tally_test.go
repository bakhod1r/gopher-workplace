package tally

import "testing"

func TestEvenStats(t *testing.T) {
	c, s := EvenStats()
	if c != 0 || s != 0 {
		t.Errorf("empty=>%d,%d want 0,0", c, s)
	}
	c, s = EvenStats(1, 2, 3, 4, 5, 6)
	if c != 3 || s != 12 {
		t.Errorf("=>%d,%d want 3,12", c, s)
	}
}
