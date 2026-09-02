# Session Store

## Intuition

A Go map is a hash table that rehashes as it grows. Two goroutines writing at once can corrupt it, so the runtime detects it and crashes the server. A mutex serialises access to the whole map.

## Approach

1. `NewSessionStore` allocates the map.
2. `Save` locks, writes, unlocks.
3. `Lookup` locks, reads with comma-ok, unlocks via defer.

## Solution

```go
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
	return &SessionStore{sessions: make(map[string]int)}
}

// Save associates token with userID.
//
// Examples:
//
//	s.Save("tok1", 7); s.Lookup("tok1") => 7, true
func (s *SessionStore) Save(token string, userID int) {
	s.mu.Lock()
	s.sessions[token] = userID
	s.mu.Unlock()
}

// Lookup returns the user ID for token and whether the session exists.
//
// Examples:
//
//	s.Lookup("unknown") => 0, false
func (s *SessionStore) Lookup(token string) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id, ok := s.sessions[token]
	return id, ok
}
```

## Walkthrough

A login handler calls `Save("tok1", 7)` and takes the lock. A concurrent request calling `Lookup("tok1")` waits, then reads `7, true`. No reader ever observes the map mid-rehash.

## Pitfalls

- Returning the internal map to callers - they could read it without the lock.
- Forgetting to lock `Lookup`; a read racing a write is still a race.
- Using a value receiver, which copies the mutex.
