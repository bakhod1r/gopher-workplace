package insertoverwrite

import "testing"

func inorder(t *Tree, out *[]int) {
	if t == nil {
		return
	}
	inorder(t.Left, out)
	*out = append(*out, t.Val)
	inorder(t.Right, out)
}

func TestInsert(t *testing.T) {
	var root *Tree
	for _, v := range []int{5, 3, 8} {
		root = Insert(root, v)
	}
	var got []int
	inorder(root, &got)
	if len(got) != 3 {
		t.Fatalf("=%v want 3 nodes", got)
	}
}
