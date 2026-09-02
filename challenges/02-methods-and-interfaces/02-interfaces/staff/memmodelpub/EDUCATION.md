# Safe Publication

## Intuition

The Go memory model gives no ordering between ordinary writes seen from another goroutine. `atomic.Pointer.Store` and `Load` create the happens-before edge: everything the publisher wrote before the store is visible to any goroutine that loads that pointer.

## Approach

1. `Publish` is a single `Store` of an already complete object.
2. `Load` returns the pointer plus a `nil` check.
3. `BuildAndPublish` copies the caller's slice, fills every field, and only then publishes.
4. A published object is never mutated — new state means a new object.

## Solution

```go
func (p *Publisher) Publish(cfg *Config) { p.cur.Store(cfg) }

func (p *Publisher) Load() (*Config, bool) {
	cfg := p.cur.Load()
	return cfg, cfg != nil
}

func (p *Publisher) Ready() bool { return p.cur.Load() != nil }

func BuildAndPublish(p *Publisher, name string, version int, tags []string) {
	cp := make([]string, len(tags))
	copy(cp, tags)

	cfg := &Config{
		Name:    name,
		Version: version,
		Tags:    cp,
	}
	p.Publish(cfg)
}
```

## Walkthrough

Publishing an empty object and filling it afterwards would break the model: the store no longer precedes the field writes, so readers may legally observe `Name == ""` — exactly what the concurrent test looks for.

## Pitfalls

- `p.cur.Store(cfg)` followed by `cfg.Tags = ...` — the fields are written after publication and readers can see the gap.
- A plain `*Config` field guarded by nothing, which is a data race regardless of how careful the ordering looks.
- Publishing a pointer to a caller-owned slice, which lets the caller mutate a supposedly immutable snapshot.
