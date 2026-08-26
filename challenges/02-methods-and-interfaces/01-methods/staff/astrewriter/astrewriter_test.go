package astrewriter

import "testing"

func TestRewrite(t *testing.T) {
	root := &Node{Type: "Ident", Val: "foo"}
	root.Rewrite()
	if root.Val != "bar" {
		t.Errorf("expected bar")
	}
}
