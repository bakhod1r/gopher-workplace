package namedret

import "testing"

func TestSafeDiv(t *testing.T) {
	q, ok := SafeDiv(10, 2)
	if q != 5 || !ok {
		t.Errorf("=%d,%v want 5,true", q, ok)
	}
	q, ok = SafeDiv(1, 0)
	if q != 0 || ok {
		t.Errorf("=%d,%v want 0,false", q, ok)
	}
}
