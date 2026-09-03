# Remember The Failure

## Intuition

Caching only successes means a permanently broken key is retried on every request — the failure path becomes the hot path. Storing the error with the value fixes that.

## Approach

1. Look the key up with the comma-ok idiom.
2. Return the stored entry when present.
3. Otherwise call `Load`, store both fields, and return them.

## Solution

```go
if e, ok := c.entries[key]; ok {
	return e.v, e.err
}
v, err := c.Load(key)
if c.entries == nil {
	c.entries = make(map[string]entry)
}
c.entries[key] = entry{v: v, err: err}
return v, err
```

## Walkthrough

The failing key is loaded once; the two later calls return the stored error without touching `Load`.

## Pitfalls

- Testing `if v, ok := …; ok && err == nil`, which re-loads every failure.
- Forgetting to initialise the map, panicking on the first write.
- Caching only the value, so a cached failure looks like a cached zero.
