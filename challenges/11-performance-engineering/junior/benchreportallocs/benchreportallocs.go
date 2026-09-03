// Package benchreportallocs — Gopher Workplace challenge.
package benchreportallocs

// Report renders the per-operation memory line that -benchmem prints:
// "<bytes> B/op\t<allocs> allocs/op", with both numbers truncated toward zero.
// A non-positive iters reports zeros.
//
// Examples:
//
//	Report(2048, 8, 4) => "512 B/op\t2 allocs/op"
func Report(bytes, allocs uint64, iters int) string {
	panic("not implemented")
}
