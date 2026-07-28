// Package defersnapshot — Gopher Workplace challenge.
package defersnapshot

// Snapshot sets x=1, defers a call that adds the CURRENT x to a result, then
// sets x=100 before returning. Because defer evaluates its arguments when the
// defer statement runs (x==1), the added value is 1, not 100.
func Snapshot() (r int) {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
