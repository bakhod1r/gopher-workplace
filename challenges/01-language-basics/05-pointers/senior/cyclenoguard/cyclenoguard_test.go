package cyclenoguard

import "testing"

func TestHasCycle(t *testing.T) {
	a := &Node{Val: 1}
	b := &Node{Val: 2}
	c := &Node{Val: 3}
	a.Next = b
	b.Next = c
	if HasCycle(a) {
		t.Errorf("acyclic reported cyclic")
	}
	c.Next = a
	if !HasCycle(a) {
		t.Errorf("cycle not detected")
	}
}
