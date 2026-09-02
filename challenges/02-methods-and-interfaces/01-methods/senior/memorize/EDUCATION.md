# Memoizer

## Intuition

Memoization is a trade: memory for time. It is only correct when the wrapped
function is pure — same input, same output, no side effects — because after the
first call the function is never consulted again for that key.

## Approach

1. Look the key up with the comma-ok form.
2. On a hit, return immediately.
3. On a miss, compute, store, return.

## Solution

```go
func (m *Memoizer) Get(key string) string {
	if v, ok := m.cache[key]; ok {
		return v
	}
	v := m.fn(key)
	m.cache[key] = v
	return v
}
```

## Walkthrough

The test counts calls in a closure. `m.Get("a")` misses, so `fn` runs, `calls`
becomes 1, and `"a-val"` is stored. The second `m.Get("a")` finds the key and
returns without touching `fn`, so `calls` stays 1 — that counter is the real
assertion, not the returned string.

## Pitfalls

- **Zero-value check instead of comma-ok.** `if m.cache[key] != ""` treats a
  cached empty result as a miss, so an expensive function returning `""` is
  called every single time.
- **Forgetting to store.** The function still returns the right value, but the
  call counter keeps climbing.
- **Nil map.** `New` allocates it; constructing a `Memoizer{}` literal directly
  and writing to `cache` panics.

## Not safe for concurrent use

Two goroutines calling `Get` at once race on the map — Go's runtime detects
concurrent map writes and crashes the process. Making this safe needs a mutex,
or `sync.Once` per key; see the `onceinit` puzzle for the single-value case.
