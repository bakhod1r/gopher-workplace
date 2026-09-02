package sortstablecheckgen

import "testing"

func key(s string) int { return len(s) }

func TestIsStableByAcceptsStable(t *testing.T) {
	original := []string{"bb", "aa", "c"}
	sorted := []string{"c", "bb", "aa"}
	if !IsStableBy(sorted, original, key) {
		t.Error("IsStableBy = false, want true")
	}
}

func TestIsStableByRejectsUnstable(t *testing.T) {
	original := []string{"bb", "aa", "c"}
	unstable := []string{"c", "aa", "bb"}
	if IsStableBy(unstable, original, key) {
		t.Error("IsStableBy = true, want false (the equal-key pair was reordered)")
	}
}

func TestIsStableByRejectsUnsorted(t *testing.T) {
	original := []string{"bb", "c"}
	unsorted := []string{"bb", "c"}
	if IsStableBy(unsorted, original, key) {
		t.Error("IsStableBy = true, want false (not sorted by key)")
	}
}

func TestIsStableByEmpty(t *testing.T) {
	if !IsStableBy([]string{}, []string{}, key) {
		t.Error("IsStableBy(empty) = false, want true")
	}
}
