package popnoguard

import "testing"

func TestPop(t *testing.T) {
	q := &Queue{}
	if _, ok := q.Pop(); ok {
		t.Errorf("empty pop should be false (no panic)")
	}
	q.head = &Node{Val: 7}
	if v, ok := q.Pop(); !ok || v != 7 {
		t.Errorf("=%d,%v want 7,true", v, ok)
	}
}
