package resetnode

import "testing"

func TestReset(t *testing.T) {
	tail := &Node{Value: 5}
	n := &Node{Value: 3, Next: tail}
	Reset(n)
	if n.Value != 0 || n.Next != nil {
		t.Errorf("not reset: %+v", *n)
	}
}
