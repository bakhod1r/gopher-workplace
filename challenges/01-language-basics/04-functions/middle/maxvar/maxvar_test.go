package maxvar

import "testing"

func TestMax(t *testing.T) {
	if _, ok := Max(); ok {
		t.Errorf("no args: ok should be false")
	}
	m, ok := Max(3, 9, 2, 9, 1)
	if m != 9 || !ok {
		t.Errorf("=%d,%v want 9,true", m, ok)
	}
}
