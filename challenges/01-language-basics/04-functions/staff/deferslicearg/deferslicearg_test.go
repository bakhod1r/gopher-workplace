package deferslicearg

import "testing"

func TestBuildAndReport(t *testing.T) {
	if got := BuildAndReport(4); got != 4 {
		t.Errorf("=%d want 4 (defer snapshotted empty header?)", got)
	}
	if got := BuildAndReport(0); got != 0 {
		t.Errorf("=%d want 0", got)
	}
}
