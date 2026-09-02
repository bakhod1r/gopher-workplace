// Package auditlog - Gopher Workplace challenge.
package auditlog

import "sync"

// AuditLog buffers compliance entries before they are flushed.
type AuditLog struct {
	mu      sync.Mutex
	entries []string
}

// Append buffers one audit entry.
//
// Examples:
//
//	var l AuditLog; l.Append("login"); l.Len() => 1
func (l *AuditLog) Append(entry string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Entries returns a copy of the buffered entries, in order.
//
// Examples:
//
//	l.Append("login"); l.Append("logout"); l.Entries() => ["login", "logout"]
func (l *AuditLog) Entries() []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Len returns the number of buffered entries.
//
// Examples:
//
//	var l AuditLog; l.Len() => 0
func (l *AuditLog) Len() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
