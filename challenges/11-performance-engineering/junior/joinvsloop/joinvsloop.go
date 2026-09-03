// Package joinvsloop — Gopher Workplace challenge.
package joinvsloop

// JoinPath joins the segments with "/" using the standard library's
// single-allocation join rather than a concatenation loop. No segments
// yields "".
//
// Examples:
//
//	JoinPath([]string{"a", "b"}) => "a/b"
func JoinPath(parts []string) string {
	panic("not implemented")
}

// SplitPath is the inverse: it splits on "/". The empty string yields an
// empty, non-nil slice.
//
// Examples:
//
//	SplitPath("a/b") => []string{"a", "b"}
func SplitPath(s string) []string {
	panic("not implemented")
}
