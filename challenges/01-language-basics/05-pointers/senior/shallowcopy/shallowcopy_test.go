package shallowcopy

import "testing"

func TestCopy(t *testing.T) {
	orig := &Tree{Val: 1, Left: &Tree{Val: 2}, Right: &Tree{Val: 3}}
	cp := Copy(orig)
	cp.Left.Val = 99
	if orig.Left.Val != 2 {
		t.Errorf("original mutated: %d (children shared)", orig.Left.Val)
	}
}
