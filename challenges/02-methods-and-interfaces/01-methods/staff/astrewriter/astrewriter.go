// Package astrewriter — Gopher Workplace challenge.
package astrewriter

type Node struct {
	Type        string
	Left, Right *Node
	Val         string
}

// Rewrite replaces "foo" idents with "bar".
func (n *Node) Rewrite() {
	// TODO(candidate): if Type == "Ident" and Val == "foo", set Val = "bar".
	// Recurse.
	panic("not implemented")
}
