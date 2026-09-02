// Package notifier — Gopher Workplace challenge.
package notifier

// Notifier delivers a message and reports success.
type Notifier interface {
	Notify(msg string) bool
}

// Email records what it delivered.
type Email struct {
	Sent []string
}

// Notify records the message and always succeeds.
func (e *Email) Notify(msg string) bool {
	// TODO(candidate): record msg, return true.
	panic("not implemented")
}

// Broken never delivers.
type Broken struct{}

// Notify always fails.
func (b Broken) Notify(msg string) bool {
	// TODO(candidate): return false.
	panic("not implemented")
}

// Broadcast notifies everyone and counts successes.
//
// Examples:
//
//	Broadcast([]Notifier{&Email{}, Broken{}}, "hi") => 1
func Broadcast(ns []Notifier, msg string) int {
	// TODO(candidate): notify all, count the true results.
	panic("not implemented")
}
