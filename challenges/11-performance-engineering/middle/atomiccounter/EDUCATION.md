# Atomics And The Retry Loop

## Intuition

Compare-and-swap is the primitive everything else is built from: "store this, but only if the value is still what I read". Failure just means someone else changed it, so read again and retry.

## Approach

1. `Add` and `Total` map directly onto the atomic type's methods.
2. `Observe` loops: load the current max, return if the new value is not larger, otherwise attempt the swap and retry on failure.

## Solution

```go
func (s *Stats) Add(delta int64) int64 { return s.total.Add(delta) }

func (s *Stats) Total() int64 { return s.total.Load() }

func (s *Stats) Observe(v int64) {
	for {
		cur := s.max.Load()
		if v <= cur {
			return
		}
		if s.max.CompareAndSwap(cur, v) {
			return
		}
	}
}

func (s *Stats) Max() int64 { return s.max.Load() }
```

## Walkthrough

The `v <= cur` check inside the loop, not before it, is what makes the retry correct: by the time a swap fails, another goroutine may have stored something bigger than `v`, and the next iteration must notice that and give up.

## Pitfalls

- `if v > s.max.Load() { s.max.Store(v) }`, which loses the maximum whenever two goroutines interleave between the load and the store.
- Retrying without re-reading, which spins forever on a stale expected value.
- Reaching for atomics on a hot shared word and expecting it to scale; the contention moved, it did not disappear.
