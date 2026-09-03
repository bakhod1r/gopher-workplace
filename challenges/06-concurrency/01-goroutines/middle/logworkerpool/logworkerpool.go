// Package logworkerpool — Gopher Workplace challenge.
package logworkerpool

// ShipLines ships every log line through a fixed pool of worker goroutines and
// returns one slot per line, in input order: nil when the line was accepted,
// the shipper's error when it was not. The pool size is fixed regardless of how
// many lines arrive, so a burst of a million lines still costs workers
// goroutines and workers open connections.
//
// A workers value of zero or less is treated as one.
//
// Examples:
//
//	ShipLines([]string{"ok", "bad"}, 2, ship)  => [<nil> errRejected]
//	ShipLines([]string{"ok"}, 8, ship)         => [<nil>]
//	ShipLines(nil, 4, ship)                    => []
func ShipLines(lines []string, workers int, ship func(line string) error) []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
