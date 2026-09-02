// Package resizequeue — Gopher Workplace challenge.
package resizequeue

// ResizeQueue drains the upload queue with a fixed pool of worker goroutines
// fed by a jobs channel and reporting on a results channel. The resized keys
// come back sorted ascending. workers is >= 1.
//
// Examples:
//
//	ResizeQueue([]string{"a", "b"}, 2, resize)  => []string{"a-512", "b-512"}
//	ResizeQueue([]string{"z"}, 4, resize)       => []string{"z-512"}
//	ResizeQueue(nil, 3, resize)                 => nil
func ResizeQueue(uploads []string, workers int, resize func(string) string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
