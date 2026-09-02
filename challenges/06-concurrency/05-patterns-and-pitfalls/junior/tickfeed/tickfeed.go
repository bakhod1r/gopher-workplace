// Package tickfeed — Gopher Workplace challenge.
package tickfeed

// LiveTicks wraps the market data feed so the consumer can walk away: it
// forwards every tick on a new channel and closes it as soon as ticks is
// closed or done is closed, whichever happens first.
//
// Examples:
//
//	done open, ticks yields 1, 2, 3  => yields 1, 2, 3 then closes
//	done open, ticks closed empty    => closes immediately
//	done already closed              => closes with no ticks
func LiveTicks(done <-chan struct{}, ticks <-chan int) <-chan int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
