# Per-Collection Lazy Index

## Intuition

`sync.Once` does two jobs at once: it guarantees the body runs a single time, and it makes every other caller *wait* for that run. That second half is what makes it a deduplicator, not just a guard. The only design work left is scoping — one `Once` per collection instead of one for the whole registry, so a cold `users` build is not blocked behind a cold `orders` build.

## Approach

1. `NewRegistry`: `make` both maps, keep `build`.
2. `Index`: `Lock`; look the `*sync.Once` up; if absent, `new(sync.Once)` and insert it; `Unlock`.
3. Call `once.Do(func(){ idx := r.build(collection); Lock; r.ready[collection] = idx; Unlock })`.
4. `Lock`, `defer Unlock`, and return `r.ready[collection]`.

## Solution

```go
// NewRegistry returns a Registry that compiles indexes with build.
//
// Examples:
//
//	NewRegistry(func(c string) string { return c + "-idx" }) != nil => true
func NewRegistry(build func(collection string) string) *Registry {
	return &Registry{
		builds: make(map[string]*sync.Once),
		ready:  make(map[string]string),
		build:  build,
	}
}

// Index returns the compiled index for a collection, building it on first use.
// build runs exactly once per collection no matter how many goroutines ask.
//
// Examples:
//
//	r := NewRegistry(f); r.Index("orders")               => "orders-idx"
//	r.Index("orders"); r.Index("orders")                 => build ran once
func (r *Registry) Index(collection string) string {
	r.mu.Lock()
	once, ok := r.builds[collection]
	if !ok {
		once = new(sync.Once)
		r.builds[collection] = once
	}
	r.mu.Unlock()

	once.Do(func() {
		idx := r.build(collection)
		r.mu.Lock()
		r.ready[collection] = idx
		r.mu.Unlock()
	})

	r.mu.Lock()
	defer r.mu.Unlock()
	return r.ready[collection]
}
```

## Walkthrough

- Thirty-two goroutines call `Index("orders")`. They serialise briefly on the mutex; the first inserts the `Once`, the rest find it.
- All thirty-two then call `Do` on the *same* `Once`: one runs `build`, the other thirty-one block inside `Do` until it returns.
- When `Do` returns, the result is already in `ready`, so every caller reads the identical index.
- `Index("users")` uses a different `Once` and proceeds in parallel with the `orders` build.

## Pitfalls

- Holding `r.mu` across `once.Do` — the build then serialises every collection, and a build that itself calls `Index` deadlocks.
- Storing `sync.Once` by value in the map: `r.builds[c]` returns a *copy*, and a copied `Once` has forgotten it already ran, so `build` runs again. Store `*sync.Once`.
- Copying a `sync.Once` (or the whole `Registry`) after first use; `go vet` flags this as copying a lock.
- Checking `ready[collection]` first and only then falling back to `Once` — two goroutines can both see it missing, and one of the two `build` calls is wasted.
