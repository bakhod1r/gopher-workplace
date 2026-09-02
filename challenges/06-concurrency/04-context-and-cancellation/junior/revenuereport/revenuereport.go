// Package revenuereport — Gopher Workplace challenge.
package revenuereport

import "context"

// TotalRevenue sums the cent amounts streamed by the report query until the
// stream closes, and returns the total. If the user cancels the export or the
// export budget expires first, it stops immediately and returns the partial
// total together with ctx.Err().
//
// Examples:
//
//	TotalRevenue(live ctx, closed chan 100,250)  => 350, nil
//	TotalRevenue(live ctx, closed empty chan)    => 0, nil
//	TotalRevenue(cancelled ctx, empty chan)      => 0, context.Canceled
func TotalRevenue(ctx context.Context, rows <-chan int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
