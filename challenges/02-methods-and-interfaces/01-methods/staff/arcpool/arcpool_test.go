package arcpool

import "testing"

func TestARC(t *testing.T) {
	a := &ARC{T1: make(map[int]bool), T2: make(map[int]bool)}

	a.Access(1)
	if !a.T1[1] || a.T2[1] {
		t.Error("1 should be in T1")
	}

	a.Access(1)
	if a.T1[1] || !a.T2[1] {
		t.Error("1 should move to T2")
	}

	a.Access(1)
	if a.T1[1] || !a.T2[1] {
		t.Error("1 should stay in T2")
	}
}
