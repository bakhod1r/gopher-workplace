// Package allocsperop — Gopher Workplace challenge.
package allocsperop

// AllocsOf reports how many heap allocations one call to f performs, averaged
// over runs samples and rounded to the nearest whole allocation. Use the
// testing package's own measurement rather than reading memory stats by hand.
// A non-positive runs is treated as 1.
//
// Examples:
//
//	AllocsOf(100, func() { _ = make([]byte, 64) }) => 1
func AllocsOf(runs int, f func()) int {
	panic("not implemented")
}
