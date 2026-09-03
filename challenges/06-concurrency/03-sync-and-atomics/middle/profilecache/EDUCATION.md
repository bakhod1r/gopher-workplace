# Read-Mostly Profile Cache

## Intuition

An `RWMutex` pays off exactly when reads dominate and the read path is short. The trap is a read path that mutates something — here, the hit counter. Promoting `Get` to a full `Lock` to bump an `int64` would throw away all the parallelism the `RWMutex` bought you. Making the counter an `atomic.Int64` keeps the read path read-only as far as the mutex is concerned.

## Approach

1. `Put`: `Lock`, assign, `Unlock`.
2. `Get`: `RLock`, read the map into locals, `RUnlock`, then `hits.Add(1)` or `misses.Add(1)`, and return.
3. `Invalidate`: `Lock` with `defer Unlock`, check presence with the comma-ok form, `delete`, return the flag.
4. `Stats`: return `hits.Load(), misses.Load()`.

## Solution

```go
// Put stores or replaces a user's profile.
//
// Examples:
//
//	c := NewCache(); c.Put("u1", "ada"); c.Get("u1") => "ada", true
func (c *Cache) Put(userID, profile string) {
	c.mu.Lock()
	c.profiles[userID] = profile
	c.mu.Unlock()
}

// Get returns a cached profile and records a hit or a miss.
//
// Examples:
//
//	c.Put("u1", "ada"); c.Get("u1")  => "ada", true
//	NewCache().Get("nobody")         => "", false
func (c *Cache) Get(userID string) (string, bool) {
	c.mu.RLock()
	profile, ok := c.profiles[userID]
	c.mu.RUnlock()
	if ok {
		c.hits.Add(1)
	} else {
		c.misses.Add(1)
	}
	return profile, ok
}

// Invalidate drops a user's entry and reports whether one was present.
//
// Examples:
//
//	c.Put("u1", "ada"); c.Invalidate("u1") => true
//	NewCache().Invalidate("u1")            => false
func (c *Cache) Invalidate(userID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.profiles[userID]
	delete(c.profiles, userID)
	return ok
}

// Stats returns the hit and miss totals observed so far.
//
// Examples:
//
//	c.Put("u1", "ada"); c.Get("u1"); c.Get("u2"); c.Stats() => 1, 1
func (c *Cache) Stats() (hits, misses int64) {
	return c.hits.Load(), c.misses.Load()
}
```

## Walkthrough

- Eight readers call `Get` concurrently: all eight hold `RLock` at once and their `Add(1)` calls serialise only inside the atomic.
- A writer calling `Put` waits for the current readers to drain, then holds the lock alone.
- `Invalidate` on an absent key deletes nothing — `delete` on a missing key is a no-op — and returns `false`.
- The concurrency test asserts `hits+misses` equals the exact number of `Get` calls, so a dropped or double-counted increment fails deterministically.

## Pitfalls

- Using `Lock` in `Get` — correct, but it serialises every reader and defeats the whole design.
- Incrementing a plain `int64` field under `RLock`: multiple readers hold that lock simultaneously, so it is a genuine data race.
- Calling `RUnlock` on a path that returns early; use locals plus `defer`, or unlock on every branch.
- Reporting `Len` from an unlocked `len(c.profiles)` — that races with `Put`.
