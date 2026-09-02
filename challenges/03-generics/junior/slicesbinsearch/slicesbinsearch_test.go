package slicesbinsearch

import "testing"

func TestFind(t *testing.T) {
	s := []int{1, 3, 5}
	if i, ok := Find(s, 3); i != 1 || !ok {
		t.Errorf("Find(s, 3) = %d, %v, want 1, true", i, ok)
	}
	if i, ok := Find(s, 1); i != 0 || !ok {
		t.Errorf("Find(s, 1) = %d, %v, want 0, true", i, ok)
	}
	if i, ok := Find(s, 4); i != 2 || ok {
		t.Errorf("Find(s, 4) = %d, %v, want 2, false (insertion point)", i, ok)
	}
	if i, ok := Find(s, 9); i != 3 || ok {
		t.Errorf("Find(s, 9) = %d, %v, want 3, false", i, ok)
	}
	if i, ok := Find([]int{}, 1); i != 0 || ok {
		t.Errorf("Find(empty, 1) = %d, %v, want 0, false", i, ok)
	}
}

func TestFindStrings(t *testing.T) {
	if i, ok := Find([]string{"a", "c"}, "c"); i != 1 || !ok {
		t.Errorf(`Find = %d, %v, want 1, true`, i, ok)
	}
}
