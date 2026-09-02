package slicesbsearchfunc

import "testing"

func TestFindByID(t *testing.T) {
	rows := []Row{{1, "a"}, {3, "b"}, {5, "c"}}
	if i, ok := FindByID(rows, 3); i != 1 || !ok {
		t.Errorf("FindByID(3) = %d, %v, want 1, true", i, ok)
	}
	if i, ok := FindByID(rows, 1); i != 0 || !ok {
		t.Errorf("FindByID(1) = %d, %v, want 0, true", i, ok)
	}
	if i, ok := FindByID(rows, 5); i != 2 || !ok {
		t.Errorf("FindByID(5) = %d, %v, want 2, true", i, ok)
	}
}

func TestFindByIDMissing(t *testing.T) {
	rows := []Row{{1, "a"}, {3, "b"}, {5, "c"}}
	if i, ok := FindByID(rows, 4); i != 2 || ok {
		t.Errorf("FindByID(4) = %d, %v, want 2, false", i, ok)
	}
	if i, ok := FindByID(rows, 9); i != 3 || ok {
		t.Errorf("FindByID(9) = %d, %v, want 3, false", i, ok)
	}
	if i, ok := FindByID(nil, 1); i != 0 || ok {
		t.Errorf("FindByID(nil) = %d, %v, want 0, false", i, ok)
	}
}
