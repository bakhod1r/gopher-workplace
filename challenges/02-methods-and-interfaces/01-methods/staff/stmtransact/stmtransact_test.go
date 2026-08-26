package stmtransact

import "testing"

func TestSTM(t *testing.T) {
	tv := &TVar{val: 5}
	Tx(tv, func(v int) int { return v * 2 })
	if tv.val != 10 {
		t.Errorf("got %d", tv.val)
	}
}
