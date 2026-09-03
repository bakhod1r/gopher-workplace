// Package nilout — Gopher Workplace challenge.
package nilout

// Node is one payload the slice points at.
type Node struct {
	ID int
}

// DropAll clears every element of s to nil, in place.
//
// The length of s must not change; only the pointers it holds are released
// so the nodes they referenced become unreachable.
//
// Examples:
//
//	s := []*Node{{1}}; DropAll(s) => s[0] == nil
func DropAll(s []*Node) {
	panic("not implemented")
}
