// Package deadletters — Gopher Workplace challenge.
package deadletters

// DrainDeadLetters consumes the dead-letter queue until it is closed and
// returns how many messages it discarded.
//
// Examples:
//
//	DrainDeadLetters(chan of 3 messages)  => 3
//	DrainDeadLetters(chan of 1 message)   => 1
//	DrainDeadLetters(closed empty)        => 0
func DrainDeadLetters(msgs <-chan string) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
