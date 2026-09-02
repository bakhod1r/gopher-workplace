// Package batchcutoff — Gopher Workplace challenge.
package batchcutoff

import "time"

// MissedCutoff builds the batch job's context from the absolute cut-off instant
// handed down by the scheduler, waits for that context to finish, and returns
// the reason. Callers pass a cut-off that has already passed, meaning the job
// started too late and must refuse to run.
//
// Examples:
//
//	MissedCutoff(time.Now().Add(-time.Hour)) => context.DeadlineExceeded
//	MissedCutoff(time.Unix(0, 0))            => context.DeadlineExceeded
//	errors.Is(..., context.Canceled)         => false
func MissedCutoff(cutoff time.Time) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
