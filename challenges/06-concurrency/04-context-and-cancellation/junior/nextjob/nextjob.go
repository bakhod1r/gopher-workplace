// Package nextjob — Gopher Workplace challenge.
package nextjob

import (
	"context"
	"errors"
)

// ErrQueueClosed reports that the job queue was closed and drained, so this
// worker should exit normally rather than treat it as a failure.
var ErrQueueClosed = errors.New("job queue closed")

// NextJob takes the next job ID off the worker pool's queue. It stops early
// when the worker's context finishes during a rolling deploy, and reports
// ErrQueueClosed once the producer has closed the queue.
//
// Examples:
//
//	NextJob(live ctx, chan with "job-1")  => "job-1", nil
//	NextJob(live ctx, closed chan)        => "", ErrQueueClosed
//	NextJob(cancelled ctx, empty chan)    => "", context.Canceled
func NextJob(ctx context.Context, jobs <-chan string) (string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
