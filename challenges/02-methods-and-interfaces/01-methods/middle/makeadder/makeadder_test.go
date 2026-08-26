package makeadder

import "testing"

func TestAdder(t *testing.T) {
	n := Number{Val: 5}
	add := n.Adder()

	if got := add(3); got != 8 {
		t.Errorf("add(3) = %d, want 8", got)
	}
	if got := add(-2); got != 3 {
		t.Errorf("add(-2) = %d, want 3", got)
	}

	// Verify n was captured by value.
	n.Val = 10
	if got := add(3); got != 8 {
		t.Errorf("closure should capture by value: add(3) = %d, want 8", got)
	}
}
