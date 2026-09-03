// Package splitviews — Gopher Workplace challenge.
package splitviews

// Fields appends each sep-separated field of line to dst as a view and
// returns the extended slice.
//
// The fields share line's storage, and dst lets the caller reuse the header
// slice, so a steady-state call allocates nothing at all.
//
// Examples:
//
//	Fields(nil, []byte("a,b"), ',') => [][]byte{"a", "b"}
func Fields(dst [][]byte, line []byte, sep byte) [][]byte {
	panic("not implemented")
}
