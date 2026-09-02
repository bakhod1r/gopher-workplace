// Package visitorifc — Gopher Workplace challenge.
package visitorifc

import "strings"

// Node is a document node.
type Node interface {
	Accept(v Visitor)
}

// Visitor performs an operation on a node.
type Visitor interface {
	Visit(n Node)
}

// Text is a leaf of words.
type Text struct {
	Content string
}

// Accept visits this node.
func (t Text) Accept(v Visitor) {
	// TODO(candidate): visit self.
	panic("not implemented")
}

// Section has a title and children.
type Section struct {
	Title    string
	Children []Node
}

// Accept visits this node, then every child in order.
func (s Section) Accept(v Visitor) {
	// TODO(candidate): visit self, then recurse into children.
	panic("not implemented")
}

// WordCounter counts words in text nodes.
type WordCounter struct {
	Words int
}

// Visit adds the word count of text nodes.
//
// Examples:
//
//	visiting Text{Content: "a b"} => Words += 2
func (w *WordCounter) Visit(n Node) {
	// TODO(candidate): count words in Text nodes only.
	panic("not implemented")
}

// HeadingCollector collects section titles in visit order.
type HeadingCollector struct {
	Titles []string
}

// Visit records section titles.
func (h *HeadingCollector) Visit(n Node) {
	// TODO(candidate): collect Section titles only.
	panic("not implemented")
}

// Walk runs v over n.
func Walk(n Node, v Visitor) {
	// TODO(candidate): start the traversal.
	panic("not implemented")
}

var _ = strings.Fields
