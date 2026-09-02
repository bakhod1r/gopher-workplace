// Package firstrow — Gopher Workplace challenge.
package firstrow

import "context"

// FirstRow returns the first row streamed by the database, or aborts if the
// request context finishes first because the client disconnected or the
// request budget expired.
//
// On abort it returns the zero row and ctx.Err().
//
// Examples:
//
//	FirstRow(live ctx, chan with "alice")   => "alice", nil
//	FirstRow(cancelled ctx, empty chan)     => "", context.Canceled
//	FirstRow(expired ctx, empty chan)       => "", context.DeadlineExceeded
func FirstRow(ctx context.Context, rows <-chan string) (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
