package clonesharetail

import "testing"

func TestCopy(t *testing.T) {
	orig := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	cp := Copy(orig)
	cp.Next.Val = 99
	if orig.Next.Val != 2 {
		t.Errorf("original mutated: %d (tail shared)", orig.Next.Val)
	}
}
