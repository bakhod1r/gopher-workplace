package zeroalloccheck

import "testing"

var sinkB []byte
var sinkI int

func TestCheckPassesForAllocationFreeWork(t *testing.T) {
	got := Check(100, 0, func() { sinkI++ })
	if got != (Result{Allocs: 0, Limit: 0, OK: true}) {
		t.Errorf("Check = %+v, want {0 0 true}", got)
	}
}

func TestCheckFailsWhenOverTheLimit(t *testing.T) {
	got := Check(100, 0, func() { sinkB = make([]byte, 64) })
	if got.Allocs != 1 {
		t.Errorf("Allocs = %d, want 1", got.Allocs)
	}
	if got.OK {
		t.Error("OK = true, want false — one allocation exceeds a limit of zero")
	}
}

func TestCheckAllowsExactlyTheLimit(t *testing.T) {
	got := Check(100, 1, func() { sinkB = make([]byte, 64) })
	if !got.OK {
		t.Errorf("Check = %+v, want OK for exactly the limit", got)
	}
}

func TestCheckClampsInputs(t *testing.T) {
	got := Check(0, -5, func() { sinkI++ })
	if got.Limit != 0 || !got.OK {
		t.Errorf("Check = %+v, want a limit of 0 and OK", got)
	}
}
