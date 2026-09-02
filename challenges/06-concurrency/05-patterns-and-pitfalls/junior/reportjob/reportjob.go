// Package reportjob — Gopher Workplace challenge.
package reportjob

import "context"

// RunReport totals the report rows concurrently. If ctx is already cancelled
// it does no work and returns 0 with the context's error.
//
// Examples:
//
//	RunReport(live ctx, []int{1, 2, 3})  => 6, nil
//	RunReport(live ctx, nil)             => 0, nil
//	RunReport(cancelled ctx, ...)        => 0, context.Canceled
func RunReport(ctx context.Context, rows []int) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
