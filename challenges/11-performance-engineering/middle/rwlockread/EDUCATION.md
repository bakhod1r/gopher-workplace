# Many Readers, One Writer

## Intuition

Readers do not interfere with each other; only writers need exclusivity. An RWMutex encodes exactly that, and the win grows with the read-to-write ratio.

## Approach

1. `Get` and `Version` take `RLock`.
2. `Replace` copies under `Lock` and increments the version.
3. `Snapshot` takes one `RLock` and reads both pieces of state.

## Solution

```go
func (c *Config) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.values[key]
	return v, ok
}

func (c *Config) Version() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.version
}

func (c *Config) Replace(values map[string]string) {
	next := make(map[string]string, len(values))
	for k, v := range values {
		next[k] = v
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values = next
	c.version++
}

func (c *Config) Snapshot() (map[string]string, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make(map[string]string, len(c.values))
	for k, v := range c.values {
		out[k] = v
	}
	return out, c.version
}
```

## Walkthrough

The copy in `Replace` happens *before* the lock is taken, so the write section is as short as two assignments — the lock is held for the swap, not for the copying.

## Pitfalls

- Storing the caller's map directly, letting them mutate live configuration with no lock at all.
- Calling `c.Version()` from `Snapshot`, which either double-locks or, worse, releases the lock between the two reads.
- Reaching for `RWMutex` on a write-heavy workload, where it is slower than a plain `Mutex`.
