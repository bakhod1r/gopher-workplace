package mirrortree

import "testing"

func TestMirror(t *testing.T) {
	root := &Tree{Val: 1,
		Left:  &Tree{Val: 2},
		Right: &Tree{Val: 3}}
	Mirror(root)
	if root.Left.Val != 3 || root.Right.Val != 2 {
		t.Errorf("not mirrored: L=%d R=%d", root.Left.Val, root.Right.Val)
	}
}
