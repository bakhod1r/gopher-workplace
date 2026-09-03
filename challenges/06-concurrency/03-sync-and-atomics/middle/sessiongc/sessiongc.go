// Package sessiongc — Gopher Workplace challenge.
package sessiongc

import "sync"

// Store holds the last-seen timestamp of every live session.
type Store struct {
	sessions sync.Map // id -> int64 last-seen tick
}

// Touch records that a session was seen at tick.
//
// Examples:
//
//	var s Store; s.Touch("u1", 10); s.Active() => 1
func (s *Store) Touch(id string, tick int64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// LastSeen returns a session's last tick and whether it is live.
//
// Examples:
//
//	s.Touch("u1", 10); s.LastSeen("u1")  => 10, true
//	var s Store; s.LastSeen("gone")      => 0, false
func (s *Store) LastSeen(id string) (int64, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Expire removes every session last seen strictly before cutoff and returns
// the removed IDs, sorted.
//
// Examples:
//
//	s.Touch("u1", 1); s.Touch("u2", 9); s.Expire(5) => ["u1"]
func (s *Store) Expire(cutoff int64) []string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Active returns how many sessions are live.
//
// Examples:
//
//	var s Store; s.Active() => 0
func (s *Store) Active() int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
