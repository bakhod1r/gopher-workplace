package clearref

import "testing"

func TestDetach(t *testing.T) {
	n := &Node{Val: 1, Next: &Node{Val: 2}}
	Detach(n)
	if n.Next != nil {
		t.Errorf("Next should be nil")
	}
	if n.Val != 1 {
		t.Errorf("Val should be unchanged")
	}
}
