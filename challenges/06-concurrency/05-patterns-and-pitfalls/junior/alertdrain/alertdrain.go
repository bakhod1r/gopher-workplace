// Package alertdrain — Gopher Workplace challenge.
package alertdrain

// CountAlerts counts the alerts consumed from the alerts channel and stops
// as soon as the alerts channel is closed or done is closed, whichever
// happens first.
//
// Examples:
//
//	done open, alerts yields 3 alerts then closes  => 3
//	done open, alerts closed with no alerts       => 0
//	done already closed                           => 0
func CountAlerts(done <-chan struct{}, alerts <-chan string) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
