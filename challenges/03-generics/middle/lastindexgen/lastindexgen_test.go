package lastindexgen

import "testing"

func TestLastIndex(t *testing.T) {
	if got := LastIndex([]int{7, 1, 7}, 7); got != 2 {
		t.Errorf("LastIndex = %d, want 2", got)
	}
	if got := LastIndex([]int{7}, 7); got != 0 {
		t.Errorf("LastIndex = %d, want 0", got)
	}
	if got := LastIndex([]int{1}, 7); got != -1 {
		t.Errorf("LastIndex = %d, want -1", got)
	}
	if got := LastIndex([]int{}, 7); got != -1 {
		t.Errorf("LastIndex(empty) = %d, want -1", got)
	}
	if got := LastIndex([]string{"a", "b", "a"}, "a"); got != 2 {
		t.Errorf("LastIndex = %d, want 2", got)
	}
}
