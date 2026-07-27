// Package ifinit — Gopher Workplace challenge.
package ifinit

// Bucket classifies a non-negative n by its remainder modulo 3, using an if
// statement with an init clause (`if r := n % 3; ...`) so the remainder is
// computed once and scoped to the if/else chain.
//
//	remainder 0 => "zero"
//	remainder 1 => "one"
//	remainder 2 => "two"
//
// Examples:
//
//	Bucket(9)  => "zero"
//	Bucket(10) => "one"
//	Bucket(11) => "two"
//	Bucket(0)  => "zero"
func Bucket(n int) string {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
