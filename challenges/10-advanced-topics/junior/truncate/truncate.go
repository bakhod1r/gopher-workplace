// Package truncate — Gopher Workplace challenge.
package truncate

// Node is one payload the slice points at.
type Node struct {
	ID int
}

// Truncate returns the first n elements of s, clearing the elements it
// drops so they no longer keep their payloads reachable.
//
// n is clamped into [0, len(s)]. The result reuses s's storage.
//
// Examples:
//
//	Truncate([]*Node{{1}, {2}}, 1) => the first element only
func Truncate(s []*Node, n int) []*Node {
	panic("not implemented")
}
