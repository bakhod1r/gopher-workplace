// Package foldcollapse — Gopher Workplace challenge.
package foldcollapse

// Collapse squashes runs of the same frame in a stack down to one, the way
// flame graph tools flatten deep recursion so the picture stays readable.
// Only *consecutive* repeats collapse: a frame that reappears after another
// frame is a genuine second entry and must survive. The input is not
// modified, and an empty stack gives an empty, non-nil result.
//
// Examples:
//
//	Collapse([]string{"a", "b", "b", "b", "c"}) => []string{"a", "b", "c"}
func Collapse(stack []string) []string {
	panic("not implemented")
}

// Depth returns how many frames the deepest run of repeats removed: the
// difference between the original length and the collapsed one.
//
// Examples:
//
//	Depth([]string{"a", "b", "b", "b"}) => 2
func Depth(stack []string) int {
	panic("not implemented")
}
