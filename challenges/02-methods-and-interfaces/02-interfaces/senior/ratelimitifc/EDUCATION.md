# Rate Limiter

## Intuition

Rate limits are time-dependent, and time is the hardest dependency to test. Injecting a `Clock` turns a flaky sleep-based test into a deterministic one, and the bucket maths stays honest.

## Approach

1. `refill` computes `elapsed / Interval` whole tokens.
2. Add them, clamp to `Burst`, and advance `last` by exactly the tokens granted — not to `now`, which would discard the remainder.
3. `Allow` refills, then consumes a token when one exists.
4. `AllowN` loops `Allow` and counts.

## Solution

```go
func (t *TokenBucket) refill() {
	if t.Interval <= 0 {
		return
	}
	now := t.clock.Now()
	elapsed := now.Sub(t.last)
	gained := int(elapsed / t.Interval)
	if gained <= 0 {
		return
	}
	t.tokens += gained
	if t.tokens > t.Burst {
		t.tokens = t.Burst
	}
	t.last = t.last.Add(time.Duration(gained) * t.Interval)
}

func (t *TokenBucket) Allow() bool {
	t.refill()
	if t.tokens <= 0 {
		return false
	}
	t.tokens--
	return true
}

func (t *TokenBucket) AllowN(n int) int {
	allowed := 0
	for i := 0; i < n; i++ {
		if t.Allow() {
			allowed++
		}
	}
	return allowed
}
```

## Walkthrough

After an hour idle, `gained` is 3600 but the clamp keeps `tokens` at 2 — so `AllowN(5)` returns 2. Advancing `last` by `gained * Interval` (not to `now`) preserves the sub-interval remainder.

## Pitfalls

- Setting `last = now` after a partial interval, which silently throws away accumulated time.
- Skipping the `Burst` clamp, so a long idle period grants an unbounded burst.
- Using `time.Now()` directly, which makes the test depend on real sleeping.
