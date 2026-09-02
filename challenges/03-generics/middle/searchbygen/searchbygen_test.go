package searchbygen

import "testing"

type row struct {
	id   int
	name string
}

func idOf(r row) int { return r.id }

func TestSearchBy(t *testing.T) {
	rows := []row{{1, "a"}, {3, "b"}, {5, "c"}}
	if i, ok := SearchBy(rows, idOf, 3); i != 1 || !ok {
		t.Errorf("SearchBy(3) = %d, %v, want 1, true", i, ok)
	}
	if i, ok := SearchBy(rows, idOf, 1); i != 0 || !ok {
		t.Errorf("SearchBy(1) = %d, %v, want 0, true", i, ok)
	}
	if i, ok := SearchBy(rows, idOf, 5); i != 2 || !ok {
		t.Errorf("SearchBy(5) = %d, %v, want 2, true", i, ok)
	}
}

func TestSearchByMissing(t *testing.T) {
	rows := []row{{1, "a"}, {3, "b"}, {5, "c"}}
	if i, ok := SearchBy(rows, idOf, 4); i != 2 || ok {
		t.Errorf("SearchBy(4) = %d, %v, want 2, false (insertion point)", i, ok)
	}
	if i, ok := SearchBy(rows, idOf, 0); i != 0 || ok {
		t.Errorf("SearchBy(0) = %d, %v, want 0, false", i, ok)
	}
	if i, ok := SearchBy(rows, idOf, 9); i != 3 || ok {
		t.Errorf("SearchBy(9) = %d, %v, want 3, false", i, ok)
	}
	if i, ok := SearchBy(nil, idOf, 1); i != 0 || ok {
		t.Errorf("SearchBy(nil) = %d, %v, want 0, false", i, ok)
	}
}

func TestSearchByFirstMatch(t *testing.T) {
	rows := []row{{1, "a"}, {1, "b"}, {2, "c"}}
	if i, ok := SearchBy(rows, idOf, 1); i != 0 || !ok {
		t.Errorf("SearchBy(1) = %d, %v, want 0, true (first match)", i, ok)
	}
}
