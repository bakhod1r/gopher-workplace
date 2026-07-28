package bstcontains

import "testing"

func TestContains(t *testing.T) {
	root := &Tree{Val: 5,
		Left:  &Tree{Val: 3, Left: &Tree{Val: 1}},
		Right: &Tree{Val: 8}}
	for _, v := range []int{5, 3, 1, 8} {
		if !Contains(root, v) {
			t.Errorf("should contain %d", v)
		}
	}
	if Contains(root, 4) {
		t.Errorf("should not contain 4")
	}
	if Contains(nil, 1) {
		t.Errorf("nil contains nothing")
	}
}
