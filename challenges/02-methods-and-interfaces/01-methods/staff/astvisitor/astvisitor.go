// Package astvisitor — Gopher Workplace challenge.
package astvisitor

type Node struct {
	Type        string // "BinOp", "Ident"
	Left, Right *Node
	Name        string
}

// Visit counts all "Ident" nodes.
func (n *Node) Visit(count *int) {
	// TODO(candidate): if nil return; if Type == "Ident" increment count;
	// recurse Left and Right.
	panic("not implemented")
}
