# Catalogue Fetch Deduplication

## Intuition

Holding one global lock across a slow upstream call would serialise the whole cache. Instead the lock guards only the map, and each entry carries its own `sync.Once` — so 200 goroutines on the same key line up behind that one `Once` while other keys proceed in parallel.

## Approach

1. `NewFetcher` allocates the map.
2. `Fetch` locks, looks the key up, creates the entry when missing, and unlocks.
3. Outside the lock, `e.once.Do` loads the value into the entry.
4. Return `e.value`.

## Solution

```go
func NewFetcher() *Fetcher {
	return &Fetcher{entries: make(map[string]*entry)}
}

func (f *Fetcher) Fetch(key string, load func(key string) string) string {
	f.mu.Lock()
	e, ok := f.entries[key]
	if !ok {
		e = &entry{}
		f.entries[key] = e
	}
	f.mu.Unlock()

	e.once.Do(func() {
		e.value = load(key)
	})
	return e.value
}
```

## Walkthrough

Two goroutines ask for `sku-1`. Both take the mutex in turn; the first installs the entry, the second finds it — so both hold the same `*entry`. Both call `e.once.Do`; exactly one runs `load`, the other blocks until it finishes and then reads the value the first one wrote.

## Pitfalls

- Calling `load` while holding the mutex: every other key blocks behind one slow upstream call.
- A `sync.Once` copied by value stops working — store `*entry`, not `entry`.
- Returning `e.value` before `Do` returns leaks the zero value to the losing goroutine.
