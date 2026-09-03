// Package scratchbuf — Gopher Workplace challenge.
package scratchbuf

// AppendJoin appends the parts to scratch, separated by sep, and returns the
// extended slice — the append-style API that lets the caller own the buffer
// and reuse it. It never allocates when scratch has room.
//
// Examples:
//
//	AppendJoin(nil, []string{"a", "b"}, "/") => []byte("a/b")
func AppendJoin(scratch []byte, parts []string, sep string) []byte {
	panic("not implemented")
}

// Sized returns how many bytes AppendJoin will append, so a caller can size
// their scratch buffer once instead of discovering the size by regrowing.
//
// Examples:
//
//	Sized([]string{"a", "b"}, "/") => 3
func Sized(parts []string, sep string) int {
	panic("not implemented")
}
