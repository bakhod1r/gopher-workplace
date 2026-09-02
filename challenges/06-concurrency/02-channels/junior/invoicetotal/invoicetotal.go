// Package invoicetotal — Gopher Workplace challenge.
package invoicetotal

// TotalCents spreads the invoice lines across workers goroutines that share
// one queue channel, and returns the invoice total in cents.
//
// workers < 1 is treated as 1.
//
// Examples:
//
//	TotalCents([]int{100, 250, 99}, 2) => 449
//	TotalCents(nil, 4)                 => 0
//	TotalCents([]int{500}, 1)          => 500
func TotalCents(lines []int, workers int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
