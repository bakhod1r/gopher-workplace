package doublecopy

import "testing"

func TestAddTax(t *testing.T) {
	p := 200
	got := AddTax(p, 10)
	if got != 220 {
		t.Errorf("AddTax(200,10)=%d want 220", got)
	}
	if p != 200 {
		t.Errorf("caller price changed: %d", p)
	}
}
