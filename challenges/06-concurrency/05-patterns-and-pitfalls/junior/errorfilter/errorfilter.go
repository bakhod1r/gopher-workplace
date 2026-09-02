// Package errorfilter — Gopher Workplace challenge.
package errorfilter

// ErrorFilter is the middle stage of the log ingest pipeline: it forwards
// only the records for which isError reports true, on a new channel that is
// closed when records is drained.
//
// Examples:
//
//	ErrorFilter(chan of "ERR a", "INFO b", hasErrPrefix)  => yields "ERR a" then closes
//	ErrorFilter(chan of only INFO records, ...)          => closes with no records
//	ErrorFilter(closed empty, ...)                       => closes immediately
func ErrorFilter(records <-chan string, isError func(string) bool) <-chan string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
