package listlen

import "testing"

func TestLength(t *testing.T) {
	if Length(nil) != 0 {
		t.Errorf("nil should be 0")
	}
	c := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	if got := Length(c); got != 3 {
		t.Errorf("=%d want 3", got)
	}
}
