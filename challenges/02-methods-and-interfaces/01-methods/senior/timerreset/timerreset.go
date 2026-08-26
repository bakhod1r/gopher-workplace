// Package timerreset — Gopher Workplace challenge.
package timerreset

import "time"

// Session expires if not pinged.
type Session struct {
	lastPing time.Time
	timeout  time.Duration
}

// New creates a session.
func New(timeout time.Duration) *Session {
	return &Session{lastPing: time.Now(), timeout: timeout}
}

// Ping resets the timer. (For testing, pass 'now' explicitly).
func (s *Session) Ping(now time.Time) {
	// TODO(candidate): update lastPing
	panic("not implemented")
}

// IsExpired returns true if timeout has elapsed since lastPing.
func (s *Session) IsExpired(now time.Time) bool {
	// TODO(candidate): check if now - lastPing > timeout
	panic("not implemented")
}
