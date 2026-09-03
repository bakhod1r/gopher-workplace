# One Miss, One Fetch, Many Waiters

## Intuition

The herd is a coordination problem, not a caching one: the work is already happening, and the other callers just need to be told when it is done. A per-key `WaitGroup` is exactly that signal.

## Approach

1. Under the lock, look the key up. If a call exists, unlock, `Wait`, return its value.
2. Otherwise create the call, `Add(1)`, register it, and unlock.
3. Run `fn`, store the result, `Done`.
4. Take the lock again to delete the entry, then return.

## Solution

```go
import "sync"

// call is one in-flight or completed fetch.
type call struct {
	wg  sync.WaitGroup
	val int
}

// Group deduplicates concurrent calls by key.
type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

// Do runs fn for key and returns its result, sharing one in-flight call
// among every concurrent caller for that key.
//
// A thundering herd on a cold cache must not become N identical fetches,
// each allocating its own result.
//
// Examples:
//
// 	g.Do("a", expensive) from 32 goroutines => expensive runs once
func (g *Group) Do(key string, fn func() int) int {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	c.val = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val
}
```

## Walkthrough

Of 32 callers, one finds no entry and becomes the owner; the other 31 find it, release the lock and block in `Wait`. When the owner finishes, all 32 return the same value and the entry is removed.

## Pitfalls

- Holding `g.mu` across `fn()`, which serialises every key and defeats the purpose.
- Writing `c.val` after `Done`, so waiters can read it before it is set.
- Never deleting the entry, which silently turns this into a cache that never expires.
