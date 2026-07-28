package derefearly

import "testing"

func TestFirstOr(t *testing.T) {
	if FirstOr(nil, 7) != 7 {
		t.Errorf("nil should give default (no panic)")
	}
	if FirstOr(&Node{Val: 3}, 7) != 3 {
		t.Errorf("want 3")
	}
}
