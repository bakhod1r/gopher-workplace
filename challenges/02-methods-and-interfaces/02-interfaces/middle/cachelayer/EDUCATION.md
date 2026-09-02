# Cache Layer

## Intuition

Because the cache implements `Source`, it can be dropped in anywhere the raw source was used — and it can even wrap another cache. The subtlety is that a cache entry must record *whether* the key exists, not just its value.

## Approach

1. `SlowSource.Get` increments `Calls`, then does a comma-ok map read.
2. `Cache.Get` first checks `c.entries` with comma-ok.
3. On a miss, call the inner source once and store an `entry` with both the value and the found flag.
4. Return the stored pair on every later call.

## Solution

```go
func (s *SlowSource) Get(key string) (string, bool) {
	s.Calls++
	v, ok := s.Data[key]
	return v, ok
}

func (c *Cache) Get(key string) (string, bool) {
	if e, ok := c.entries[key]; ok {
		return e.value, e.found
	}
	v, found := c.inner.Get(key)
	c.entries[key] = entry{value: v, found: found}
	return v, found
}
```

## Walkthrough

Caching `map[string]string` alone cannot distinguish "absent" from "present but empty": the empty-value test would then hit the source twice. Storing `entry{value, found}` keyed by presence in `c.entries` fixes both cases.

## Pitfalls

- Using `map[string]string` and treating `""` as a miss — empty values and missing keys become indistinguishable.
- Not caching negative lookups, so every miss reaches the source.
- Counting the call after the map read in `SlowSource` — same result here, but count what the method actually served.
