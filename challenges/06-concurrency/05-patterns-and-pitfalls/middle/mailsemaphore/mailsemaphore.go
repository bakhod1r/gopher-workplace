// Package mailsemaphore — Gopher Workplace challenge.
package mailsemaphore

// SendAll delivers every message, running at most limit sends at the same
// time, and returns each message's provider response in input order. A limit
// of zero or less means unlimited.
//
// Examples:
//
//	SendAll([]string{"m1","m2"}, 1, send)  => ["sent:m1" "sent:m2"]
//	SendAll([]string{"m1"}, 5, send)       => ["sent:m1"]
//	SendAll(nil, 2, send)                  => []
func SendAll(messages []string, limit int, send func(msg string) string) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}
