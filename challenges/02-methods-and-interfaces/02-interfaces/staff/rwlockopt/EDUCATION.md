# Read-Write Lock

## Intuition

An `RWMutex` encodes what a plain mutex cannot express: reads do not conflict with each other. On read-heavy data that turns a serialisation point into genuine parallelism — at the cost of a heavier lock when writes are frequent.

## Approach

1. `RWStore.Get` and `Len` take `RLock`; `Set` takes `Lock`.
2. `Snapshot` copies the map under `RLock`, so no writer can interleave mid-copy.
3. `MutexStore` does the same operations under one exclusive lock.
4. The reader-tracking helpers make the overlap observable in a test.

## Solution

```go
func (s *RWStore) Get(key string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	s.enterRead()
	defer s.exitRead()

	v, ok := s.data[key]
	return v, ok
}

func (s *RWStore) Set(key string, v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = v
}

func (s *RWStore) Snapshot() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make(map[string]int, len(s.data))
	for k, v := range s.data {
		out[k] = v
	}
	return out
}
```

## Walkthrough

`TestReadersOverlap` runs eight readers and records the peak count inside the critical section. Under a plain `Lock` the peak would be 1; under `RLock` it climbs, which is the whole reason to pay for an `RWMutex`.

## Pitfalls

- Calling `Lock` while holding `RLock` — `RWMutex` is not reentrant and this deadlocks.
- Returning `s.data` from `Snapshot` instead of a copy, handing callers a map that writers mutate.
- Reaching for `RWMutex` on write-heavy data, where its extra bookkeeping loses to a plain mutex.
