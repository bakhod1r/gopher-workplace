// Package childisolation — Gopher Workplace challenge.
package childisolation

// RequestErrAfterQueryCancel builds the per-request context, derives a query
// context from it, cancels only the query (the handler abandoned that lookup
// and will fall back to the cache), and returns the error the *request*
// context reports.
//
// Examples:
//
//	RequestErrAfterQueryCancel()          => nil
//	the request context stays usable      => true
//	the result is never context.Canceled  => true
func RequestErrAfterQueryCancel() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
