// Package benchfixedinput — Gopher Workplace challenge.
package benchfixedinput

// FixedInput returns n pseudo-random values that depend only on seed, so two
// benchmark runs with the same seed measure exactly the same work.
//
// Generate each value with the 32-bit linear congruential step
//
//	state = state*1664525 + 1013904223
//
// starting from state = seed, taking the new state as the value. A
// non-positive n returns an empty, non-nil slice.
//
// Examples:
//
//	FixedInput(1, 1) => []uint32{1664526}
func FixedInput(seed uint32, n int) []uint32 {
	panic("not implemented")
}
