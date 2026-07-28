package switchinit

import "testing"

func TestClassify(t *testing.T) {
	l, r := Classify(5)
	if l != "pos" || r != 1 {
		t.Errorf("=%q,%d want pos,1", l, r)
	}
	l, r = Classify(0)
	if l != "zero" || r != 1 {
		t.Errorf("=%q,%d want zero,1", l, r)
	}
	l, r = Classify(-3)
	if l != "neg" || r != 1 {
		t.Errorf("=%q,%d want neg,1", l, r)
	}
}
