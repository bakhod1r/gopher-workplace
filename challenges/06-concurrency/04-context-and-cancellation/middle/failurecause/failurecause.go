// Package failurecause — Gopher Workplace challenge.
package failurecause

import (
	"context"
	"errors"
)

// ErrQuotaExceeded is the cause used when the export runs past its row quota.
var ErrQuotaExceeded = errors.New("export quota exceeded")

// Export scans rows until either the scan is done or the quota is exceeded,
// and returns the number of rows written plus the reason the run stopped.
//
// It cancels its own context with a cause: ErrQuotaExceeded when the quota is
// hit, nil when every row was written. The returned error is the cause, not
// the bare context.Canceled.
//
// Examples:
//
//	Export(ctx, 5 rows, quota 10) => 5, nil
//	Export(ctx, 5 rows, quota 3)  => 3, ErrQuotaExceeded
//	Export(cancelled ctx, rows)   => 0, context.Canceled
func Export(ctx context.Context, rows []string, quota int) (int, error) {
	// TODO(candidate): implement this using context.WithCancelCause and
	// context.Cause.
	panic("not implemented")
}
