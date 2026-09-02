// Package exportrows — Gopher Workplace challenge.
package exportrows

import "context"

// ExportRows renders each account ID into a CSV line, streaming the work
// through a goroutine so the consumer can abandon the export the moment the
// request context finishes.
//
// On cancellation it returns nil and ctx.Err(); otherwise the full slice.
//
// Examples:
//
//	ExportRows(live ctx, []int{1, 2})  => ["row-1", "row-2"], nil
//	ExportRows(live ctx, nil)          => [], nil
//	ExportRows(cancelled ctx, []int{1}) => nil, context.Canceled
func ExportRows(ctx context.Context, ids []int) ([]string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
