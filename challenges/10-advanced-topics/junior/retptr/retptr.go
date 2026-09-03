// Package retptr — Gopher Workplace challenge.
package retptr

// New returns a pointer to a fresh int holding v.
//
// The pointer outlives the call, so the int cannot live in the frame — the
// compiler moves it to the heap. That is one allocation, and exactly one.
//
// Examples:
//
//	*New(7) => 7
func New(v int) *int {
	panic("not implemented")
}
