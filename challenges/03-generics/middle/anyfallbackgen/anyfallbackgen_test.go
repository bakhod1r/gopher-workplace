package anyfallbackgen

import "testing"

func TestDeepCount(t *testing.T) {
	if got := DeepCount([]any{1, 2}); got != 2 {
		t.Errorf("DeepCount = %d, want 2", got)
	}
	if got := DeepCount([]any{1, []any{2, 3}}); got != 3 {
		t.Errorf("DeepCount = %d, want 3", got)
	}
	if got := DeepCount([]any{[]any{[]any{1}}}); got != 1 {
		t.Errorf("DeepCount = %d, want 1", got)
	}
}

func TestDeepCountEdges(t *testing.T) {
	if got := DeepCount(nil); got != 0 {
		t.Errorf("DeepCount(nil) = %d, want 0", got)
	}
	if got := DeepCount([]any{}); got != 0 {
		t.Errorf("DeepCount(empty) = %d, want 0", got)
	}
	if got := DeepCount(42); got != 1 {
		t.Errorf("DeepCount(leaf) = %d, want 1", got)
	}
	if got := DeepCount([]any{nil, 1}); got != 1 {
		t.Errorf("DeepCount = %d, want 1", got)
	}
}
