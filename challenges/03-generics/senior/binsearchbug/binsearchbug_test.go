package binsearchbug

import "testing"

type row struct{ id int }

func idOf(r row) int { return r.id }

func TestSearchByDuplicates(t *testing.T) {
	rows := []row{{1}, {1}, {2}}
	if i, ok := SearchBy(rows, idOf, 1); i != 0 || !ok {
		t.Errorf("SearchBy(1) = %d, %v, want 0, true (first match)", i, ok)
	}
	if i, ok := SearchBy(rows, idOf, 2); i != 2 || !ok {
		t.Errorf("SearchBy(2) = %d, %v, want 2, true", i, ok)
	}
}

func TestSearchByUnique(t *testing.T) {
	rows := []row{{1}, {3}, {5}}
	if i, ok := SearchBy(rows, idOf, 3); i != 1 || !ok {
		t.Errorf("SearchBy(3) = %d, %v, want 1, true", i, ok)
	}
	if i, ok := SearchBy(rows, idOf, 5); i != 2 || !ok {
		t.Errorf("SearchBy(5) = %d, %v, want 2, true", i, ok)
	}
}

func TestSearchByMissing(t *testing.T) {
	rows := []row{{1}, {3}, {5}}
	if i, ok := SearchBy(rows, idOf, 2); i != 1 || ok {
		t.Errorf("SearchBy(2) = %d, %v, want 1, false", i, ok)
	}
	if i, ok := SearchBy(rows, idOf, 9); i != 3 || ok {
		t.Errorf("SearchBy(9) = %d, %v, want 3, false", i, ok)
	}
	if i, ok := SearchBy(nil, idOf, 1); i != 0 || ok {
		t.Errorf("SearchBy(nil) = %d, %v, want 0, false", i, ok)
	}
}
