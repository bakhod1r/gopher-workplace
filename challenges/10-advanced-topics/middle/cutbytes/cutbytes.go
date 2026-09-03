// Package cutbytes — Gopher Workplace challenge.
package cutbytes

// Cut splits s around the first occurrence of sep.
//
// When sep is absent, before is s and after is empty. Both results are
// substrings, so nothing is copied.
//
// Examples:
//
//	Cut("a=b", '=') => "a", "b", true
func Cut(s string, sep byte) (before, after string, found bool) {
	panic("not implemented")
}
