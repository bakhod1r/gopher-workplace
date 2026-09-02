// Package auditsplit — Gopher Workplace challenge.
package auditsplit

// TeeAudit duplicates the audit event stream: every event is delivered to
// both returned channels, and both are closed when events is drained.
//
// Examples:
//
//	TeeAudit(chan of "login")   => both channels yield "login" then close
//	TeeAudit(chan of 3 events)  => both channels yield all 3 in order
//	TeeAudit(closed empty)      => both channels close immediately
func TeeAudit(events <-chan string) (<-chan string, <-chan string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
