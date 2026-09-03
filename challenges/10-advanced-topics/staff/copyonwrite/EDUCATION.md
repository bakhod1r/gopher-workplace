# Publish A New Map, Never Edit The Old One

## Intuition

When reads vastly outnumber writes, moving the cost to the writer is a good trade. The invariant that makes it safe is absolute: once a map is published, it is frozen, because you can never know who is still reading it.

## Approach

1. Take the writer mutex.
2. Load the current snapshot, copy it into a new map sized `len(old)+1`.
3. Set the key in the copy and `Store` its address.

## Solution

```go
import (
	"sync"
	"sync/atomic"
)

// Store is a read-mostly map published by pointer swap.
type Store struct {
	mu sync.Mutex // serialises writers only
	m  atomic.Pointer[map[string]int]
}

// Get reads from the current snapshot without locking.
func (s *Store) Get(key string) (int, bool) {
	p := s.m.Load()
	if p == nil {
		return 0, false
	}
	v, ok := (*p)[key]
	return v, ok
}

// Len reports the current snapshot's size.
func (s *Store) Len() int {
	p := s.m.Load()
	if p == nil {
		return 0
	}
	return len(*p)
}

// Set publishes a new snapshot of the map with key set to val.
//
// Readers hold whatever snapshot was current when they loaded it, so a
// published map must never be modified again: build a copy, then swap.
//
// Examples:
//
// 	s.Set("a", 1); s.Get("a") => 1, true
func (s *Store) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var old map[string]int
	if p := s.m.Load(); p != nil {
		old = *p
	}
	next := make(map[string]int, len(old)+1)
	for k, v := range old {
		next[k] = v
	}
	next[key] = val
	s.m.Store(&next)
}
```

## Walkthrough

A reader that loaded the old pointer keeps reading the old map, which nobody touches. The next reader loads the new pointer. There is no moment at which any map is both published and being written.

## Pitfalls

- Writing into the loaded map before swapping, which is a data race with every current reader.
- Dropping the writer lock, which lets two writers each copy the same old map and lose one update.
- Using this shape for a write-heavy map, where copying per write is far worse than a mutex.
