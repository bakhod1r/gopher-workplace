// Package ingestpipeline — Gopher Workplace challenge.
package ingestpipeline

import (
	"context"
)

// Normalise cleans one raw event record. Returning an error aborts the whole
// ingest run.
type Normalise func(ctx context.Context, record string) (string, error)

// RunIngestPipeline streams raw event records through a producer goroutine into
// a normalising stage and collects the results. Producer and consumer share one
// derived context: whether the run ends from a bad record, a cancelled upload
// or an expired budget, the same cancel unblocks the producer so it can never
// be left parked on a send nobody will receive.
//
// It returns the normalised records in input order, or nil and the reason the
// run was abandoned.
//
// Examples:
//
//	RunIngestPipeline(ctx, ["a" "b"], upper)          => ["A" "B"], nil
//	RunIngestPipeline(ctx, ["a" "bad" "c"], upper)    => nil, errBadRecord
//	RunIngestPipeline(cancelled ctx, ["a"], upper)    => nil, context.Canceled
func RunIngestPipeline(ctx context.Context, records []string, normalise Normalise) ([]string, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
