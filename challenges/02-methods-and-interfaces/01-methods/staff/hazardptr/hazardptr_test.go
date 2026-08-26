package hazardptr

import (
	"sync/atomic"
	"testing"
)

func TestHazard(t *testing.T) {
	h := &Hazard{}
	val := 42

	var shared atomic.Pointer[int]
	shared.Store(&val)

	got := h.Protect(&shared)
	if got == nil || *got != 42 {
		t.Errorf("failed to protect")
	}
}
