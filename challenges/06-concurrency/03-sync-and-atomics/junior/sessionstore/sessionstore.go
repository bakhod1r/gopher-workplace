// Package sessionstore - Gopher Workplace challenge.
package sessionstore

import "sync"

// SessionStore maps session tokens to user IDs, safely for concurrent use.
type SessionStore struct {
	mu       sync.Mutex
	sessions map[string]int
}

// NewSessionStore returns an empty, ready-to-use store.
//
// Examples:
//
//	NewSessionStore().Lookup("tok1") => 0, false
func NewSessionStore() *SessionStore {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Save associates token with userID.
//
// Examples:
//
//	s.Save("tok1", 7); s.Lookup("tok1") => 7, true
func (s *SessionStore) Save(token string, userID int) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Lookup returns the user ID for token and whether the session exists.
//
// Examples:
//
//	s.Lookup("unknown") => 0, false
func (s *SessionStore) Lookup(token string) (int, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
