package lastindexbug

import "testing"

func TestLastIndexRepeated(t *testing.T) {
	if got := LastIndex([]int{1, 2, 1}, 1); got != 2 {
		t.Errorf("LastIndex = %d, want 2", got)
	}
}

func TestLastIndexUnique(t *testing.T) {
	if got := LastIndex([]int{1, 2, 3}, 2); got != 1 {
		t.Errorf("LastIndex = %d, want 1", got)
	}
}

func TestLastIndexMissing(t *testing.T) {
	if got := LastIndex([]int{1}, 9); got != -1 {
		t.Errorf("LastIndex = %d, want -1", got)
	}
}
