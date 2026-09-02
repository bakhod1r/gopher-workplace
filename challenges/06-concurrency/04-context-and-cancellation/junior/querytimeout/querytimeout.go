// Package querytimeout — Gopher Workplace challenge.
package querytimeout

// ExhaustedQueryBudget models a database query whose time budget is already
// spent by the time the query is dispatched: the surrounding request used it
// all up. It builds the query context with a zero timeout, waits for it to
// finish, and returns the reason.
//
// The timeout must be zero or negative so the context is done immediately —
// never depend on wall-clock time passing.
//
// Examples:
//
//	ExhaustedQueryBudget()                                     => context.DeadlineExceeded
//	errors.Is(ExhaustedQueryBudget(), context.DeadlineExceeded) => true
//	errors.Is(ExhaustedQueryBudget(), context.Canceled)         => false
func ExhaustedQueryBudget() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
