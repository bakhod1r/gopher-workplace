package defersnapshot

import "testing"

func TestSnapshot(t *testing.T) {
	if got := Snapshot(); got != 1 {
		t.Errorf("=%d want 1 (defer captured x at defer-time)", got)
	}
}
