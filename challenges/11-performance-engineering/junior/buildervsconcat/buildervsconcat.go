// Package buildervsconcat — Gopher Workplace challenge.
package buildervsconcat

// JoinSep concatenates parts with sep between them, reaching the final string
// in a single allocation: compute the exact size first, then build into a
// buffer of that size. No parts yields "".
//
// Examples:
//
//	JoinSep([]string{"a", "b"}, ", ") => "a, b"
func JoinSep(parts []string, sep string) string {
	panic("not implemented")
}
