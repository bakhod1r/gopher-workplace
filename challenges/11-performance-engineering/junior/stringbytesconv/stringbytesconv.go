// Package stringbytesconv — Gopher Workplace challenge.
package stringbytesconv

// CountByte returns how many times b occurs in s. Converting the string to a
// []byte would copy the whole thing; indexing it does not, so this must not
// allocate at all.
//
// Examples:
//
//	CountByte("hello", 'l') => 2
func CountByte(s string, b byte) int {
	panic("not implemented")
}

// HasPrefixByte reports whether s starts with b, without allocating.
//
// Examples:
//
//	HasPrefixByte("go", 'g') => true
func HasPrefixByte(s string, b byte) bool {
	panic("not implemented")
}
