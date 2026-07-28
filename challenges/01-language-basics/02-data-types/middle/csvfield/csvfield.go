// Package csvfield quotes a single CSV field per RFC 4180.
package csvfield

// Quote returns s as a CSV field: if s contains a comma, double-quote, or
// newline, it is wrapped in double quotes with inner quotes doubled; otherwise
// s is returned unchanged.
//
// TODO(candidate): implement.
func Quote(s string) string {
	panic("not implemented")
}
