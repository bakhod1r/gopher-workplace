package averagef

import "testing"

func TestAverage(t *testing.T) {
	if _, ok := Average(); ok {
		t.Errorf("no args: ok should be false")
	}
	m, ok := Average(2, 4)
	if m != 3 || !ok {
		t.Errorf("Average(2,4)=%v,%v want 3,true", m, ok)
	}
	m, ok = Average(1, 2, 3, 4)
	if m != 2.5 || !ok {
		t.Errorf("=%v,%v want 2.5,true", m, ok)
	}
}
