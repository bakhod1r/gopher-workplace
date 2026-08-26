package astvisitor

import "testing"

func TestVisit(t *testing.T) {
	root := &Node{
		Type: "BinOp",
		Left: &Node{Type: "Ident", Name: "x"},
		Right: &Node{
			Type:  "BinOp",
			Left:  &Node{Type: "Ident", Name: "y"},
			Right: &Node{Type: "Num"},
		},
	}

	count := 0
	root.Visit(&count)
	if count != 2 {
		t.Errorf("got %d, want 2", count)
	}
}
