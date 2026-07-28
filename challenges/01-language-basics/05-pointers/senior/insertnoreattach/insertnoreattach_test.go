package insertnoreattach

import "testing"

func TestInsert(t *testing.T) {
	root := &Tree{Val: 5}
	root = Insert(root, 3)
	root = Insert(root, 2)
	if root.Left == nil || root.Left.Val != 3 {
		t.Fatalf("3 not inserted on the left")
	}
	if root.Left.Left == nil || root.Left.Left.Val != 2 {
		t.Fatalf("2 not inserted (return discarded)")
	}
}
