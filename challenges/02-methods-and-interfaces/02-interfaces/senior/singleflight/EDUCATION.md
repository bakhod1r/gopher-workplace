# Single Flight

## Intuition

Duplicate suppression is not caching: it only collapses calls that overlap in time. The leader does the work and the followers wait on a `WaitGroup` that acts as a one-shot broadcast.

## Approach

1. Lock the map; if a call for the key exists, unlock, `wg.Wait()`, and return its value.
2. Otherwise create a `call`, `wg.Add(1)`, register it, and unlock *before* loading.
3. Run the loader, store the result, and `wg.Done()` to release every waiter.
4. Delete the entry under the lock so the next call starts a fresh flight.

## Solution

```go
func (g *Group) Do(key string, l Loader) string {
	g.mu.Lock()
	if c, ok := g.calls[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val
	}

	c := &call{}
	c.wg.Add(1)
	g.calls[key] = c
	g.mu.Unlock()

	c.val = l.Load(key)
	c.wg.Done()

	g.mu.Lock()
	delete(g.calls, key)
	g.mu.Unlock()

	return c.val
}
```

## Walkthrough

Writing `c.val` before `wg.Done()` and reading it after `wg.Wait()` is what makes the shared value race-free: the `WaitGroup` establishes the happens-before edge between the leader's write and every follower's read.

## Pitfalls

- Holding `g.mu` across `l.Load` — every key serialises and distinct keys block each other.
- Reading `c.val` before `wg.Wait()` returns, which is a genuine data race.
- Never deleting the entry, which turns the group into a permanent cache of stale values.
