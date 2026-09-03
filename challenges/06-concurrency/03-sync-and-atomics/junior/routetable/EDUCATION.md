# Gateway Routing Table

## Intuition

A routing table is read on every request and written once a deploy. That is exactly the shape `RWMutex` is built for. The subtle part is `Snapshot`: handing back the live map would leak unsynchronised access past the lock, so copy it while still holding the lock.

## Approach

1. `NewTable` allocates the map — a nil map panics on write.
2. `Set` takes the write lock and assigns.
3. `Lookup` takes the read lock and returns the comma-ok result.
4. `Snapshot` takes the read lock, allocates a map of the right size, copies every entry, returns it.

## Solution

```go
func NewTable() *Table {
	return &Table{routes: make(map[string]string)}
}

func (t *Table) Set(path, backend string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.routes[path] = backend
}

func (t *Table) Lookup(path string) (string, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	backend, ok := t.routes[path]
	return backend, ok
}

func (t *Table) Snapshot() map[string]string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	out := make(map[string]string, len(t.routes))
	for path, backend := range t.routes {
		out[path] = backend
	}
	return out
}
```

## Walkthrough

Fifty request goroutines call `Lookup` and hold `RLock` simultaneously — none blocks another. A deploy goroutine calls `Set`, which waits for the readers to leave, then holds the table exclusively for one assignment. `Snapshot` copies under `RLock`, so a later `Set` cannot change what the caller already received.

## Pitfalls

- Returning `t.routes` directly from `Snapshot` — the caller then reads the map while a writer mutates it, and `-race` reports it.
- Using `RLock` around a write: read locks do not exclude each other, so the map corrupts.
- Forgetting `make` in `NewTable`: assignment to a nil map panics.
