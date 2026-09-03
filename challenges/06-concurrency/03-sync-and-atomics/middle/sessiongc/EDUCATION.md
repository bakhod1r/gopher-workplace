# Session Expiry Sweeper

## Intuition

The sweeper and the request path fight over the same table. A `sync.Map` lets each of them work key by key: the sweeper walks and deletes, the requests keep storing, and neither ever holds a lock over the whole table.

## Approach

1. `Touch`: `s.sessions.Store(id, tick)`.
2. `LastSeen`: `Load`, type-assert to `int64`, return the comma-ok result.
3. `Expire`: `Range`, and for each entry older than the cutoff call `LoadAndDelete`, collecting the IDs that were actually removed.
4. Sort the removed IDs before returning; `Active` counts with a `Range`.

## Solution

```go
import (
	"sort"
	"sync"
)

func (s *Store) Touch(id string, tick int64) {
	s.sessions.Store(id, tick)
}

func (s *Store) LastSeen(id string) (int64, bool) {
	v, ok := s.sessions.Load(id)
	if !ok {
		return 0, false
	}
	return v.(int64), true
}

func (s *Store) Expire(cutoff int64) []string {
	removed := []string{}
	s.sessions.Range(func(k, v any) bool {
		if v.(int64) < cutoff {
			if _, loaded := s.sessions.LoadAndDelete(k); loaded {
				removed = append(removed, k.(string))
			}
		}
		return true
	})
	sort.Strings(removed)
	return removed
}

func (s *Store) Active() int {
	n := 0
	s.sessions.Range(func(_, _ any) bool {
		n++
		return true
	})
	return n
}
```

## Walkthrough

With `u1@1` and `u2@9` and cutoff 5, `Range` visits both entries. `u1` is older, so `LoadAndDelete` removes it and reports `loaded == true`, and the ID joins the result. `u2` stays. Meanwhile a request goroutine storing `fresh@100` is unaffected — it never waits for the sweep.

## Pitfalls

- `Delete` instead of `LoadAndDelete` loses the "was it really there" answer and can double-report under two sweepers.
- Treating `Range` order as stable — always sort what you return.
- Using `<=` for the cutoff, which expires a session seen exactly at the boundary.
