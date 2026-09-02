package divmodgen

import "testing"

func TestDivMod(t *testing.T) {
	q, r, ok := DivMod(7, 2)
	if q != 3 || r != 1 || !ok {
		t.Errorf("DivMod(7, 2) = %v, %v, %v, want 3, 1, true", q, r, ok)
	}
	q, r, ok = DivMod(6, 3)
	if q != 2 || r != 0 || !ok {
		t.Errorf("DivMod(6, 3) = %v, %v, %v, want 2, 0, true", q, r, ok)
	}
}

func TestDivModNegatives(t *testing.T) {
	q, r, _ := DivMod(-7, 2)
	if q != -3 || r != -1 {
		t.Errorf("DivMod(-7, 2) = %v, %v, want -3, -1 (Go truncates)", q, r)
	}
	q, r, _ = DivMod(7, -2)
	if q != -3 || r != 1 {
		t.Errorf("DivMod(7, -2) = %v, %v, want -3, 1", q, r)
	}
}

func TestDivModByZero(t *testing.T) {
	q, r, ok := DivMod(1, 0)
	if q != 0 || r != 0 || ok {
		t.Errorf("DivMod(1, 0) = %v, %v, %v, want 0, 0, false (must not panic)", q, r, ok)
	}
}
