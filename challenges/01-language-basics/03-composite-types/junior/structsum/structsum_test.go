package structsum

import "testing"

func TestTotal(t *testing.T) {
	orders := []Order{
		{"pen", 150, 2},
		{"pad", 300, 1},
		{"ink", 500, 3},
	}
	if got := Total(orders); got != 2100 {
		t.Errorf("Total=%d; want 2100", got)
	}
	if Total(nil) != 0 {
		t.Error("empty total should be 0")
	}
}
