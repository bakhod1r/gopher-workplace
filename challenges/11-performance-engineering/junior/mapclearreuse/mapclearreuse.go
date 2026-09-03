// Package mapclearreuse — Gopher Workplace challenge.
package mapclearreuse

// Reset empties m in place, keeping the same map — and its already-allocated
// buckets — so the caller's other references see the change and the next
// round of inserts does not reallocate. A nil map is left alone.
//
// Examples:
//
//	Reset(map[string]int{"a": 1}) leaves an empty map with the same identity
func Reset(m map[string]int) {
	panic("not implemented")
}

// Tally counts the words into m after emptying it, and returns m.
//
// Examples:
//
//	Tally(m, []string{"a", "a", "b"}) => {"a":2, "b":1}
func Tally(m map[string]int, words []string) map[string]int {
	panic("not implemented")
}
