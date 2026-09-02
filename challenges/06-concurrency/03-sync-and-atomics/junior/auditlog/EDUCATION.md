# Audit Log Buffer

## Intuition

Returning `l.entries` hands the caller a window onto live storage: the next `Append` can reallocate or overwrite what the flusher is reading. Copying inside the critical section gives the caller a private snapshot that no lock is needed to read.

## Approach

1. `Append` locks, appends, unlocks.
2. `Entries` locks, allocates a slice of the right length, `copy`s into it, and returns it.
3. `Len` locks and returns `len(l.entries)`.

## Solution

```go
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
	l.mu.Lock()
	l.entries = append(l.entries, entry)
	l.mu.Unlock()
}

// Entries returns a copy of the buffered entries, in order.
//
// Examples:
//
//	l.Append("login"); l.Append("logout"); l.Entries() => ["login", "logout"]
func (l *AuditLog) Entries() []string {
	l.mu.Lock()
	defer l.mu.Unlock()
	out := make([]string, len(l.entries))
	copy(out, l.entries)
	return out
}

// Len returns the number of buffered entries.
//
// Examples:
//
//	var l AuditLog; l.Len() => 0
func (l *AuditLog) Len() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	return len(l.entries)
}
```

## Walkthrough

The flusher calls `Entries` and gets its own array of two strings. A handler then appends a third entry; the flusher's snapshot is untouched, and nothing races.

## Pitfalls

- Returning `l.entries` directly, so the caller reads shared storage without the lock.
- Returning `l.entries[:]`, which is the same aliasing bug in disguise.
- Appending outside the lock because "append is just one line".
