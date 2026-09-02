package counterstruct

import "testing"

func TestCounter(t *testing.T) {
	c := NewCounter[string]()
	if c.Count("never") != 0 {
		t.Errorf(`Count("never") = %d, want 0`, c.Count("never"))
	}
	c.Inc("a")
	c.Inc("a")
	c.Inc("b")
	if got := c.Count("a"); got != 2 {
		t.Errorf(`Count("a") = %d, want 2`, got)
	}
	if got := c.Count("b"); got != 1 {
		t.Errorf(`Count("b") = %d, want 1`, got)
	}
	if got := c.Total(); got != 3 {
		t.Errorf("Total() = %d, want 3", got)
	}
}

func TestCounterInts(t *testing.T) {
	c := NewCounter[int]()
	c.Inc(7)
	if c.Count(7) != 1 || c.Total() != 1 {
		t.Error("int counter is wrong")
	}
}
