package bstinsert

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
	for _, v := range []int{5, 3, 8, 1, 4} {
		root = Insert(root, v)
	}
	var got []int
	inorder(root, &got)
	want := []int{1, 3, 4, 5, 8}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("=%v want %v", got, want)
		}
	}
}
