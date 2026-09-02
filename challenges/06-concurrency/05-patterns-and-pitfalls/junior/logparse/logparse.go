// Package logparse — Gopher Workplace challenge.
package logparse

// ParseStage is the first stage of the log ingest pipeline: it reads raw log
// lines, applies parse to each one, and forwards the parsed record on a new
// channel that is closed when lines is drained.
//
// Examples:
//
//	ParseStage(chan of "a", "b", strings.ToUpper)  => yields "A", "B" then closes
//	ParseStage(chan of "x", identity)             => yields "x" then closes
//	ParseStage(closed empty, parse)               => closes immediately
func ParseStage(lines <-chan string, parse func(string) string) <-chan string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
