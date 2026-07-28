// Package clamp restricts a value to a range using block-scoped locals.
package clamp

// Clamp returns v limited to [lo, hi]. If lo > hi the bounds are swapped first.
// Use a short variable declaration in an if-init to normalize the bounds
// without shadowing lo/hi in the outer scope.
//
// TODO(candidate): implement.
func Clamp(v, lo, hi int) int {
	panic("not implemented")
}
