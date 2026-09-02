# ARC List Promotion

## Intuition

ARC's insight is that recency and frequency deserve separate lists. A key that
has been touched once is *recent* (T1); a key touched again is *frequent* (T2)
and should be far harder to evict. `Access` is the state machine that moves keys
between those two populations.

## Approach

1. Promotion first: if it is in T1, move it.
2. Already frequent: nothing to do.
3. Otherwise it is new: admit it to T1.

## Solution

```go
func (a *ARC) Access(key int) {
	if a.T1[key] {
		delete(a.T1, key)
		a.T2[key] = true
		return
	}
	if a.T2[key] {
		return
	}
	a.T1[key] = true
}
```

## Walkthrough

- First `Access(1)`: absent from both, so it lands in T1.
- Second: found in T1, so it is deleted there and inserted into T2. The test
  asserts *both* halves — `!a.T1[1]` and `a.T2[1]` — because forgetting the
  delete leaves the key in two lists at once.
- Third: found in T2, early return, no change.

## Pitfalls

- **Setting `a.T1[key] = false` instead of deleting.** The membership test still
  reads false, so the assertions pass — but the map keeps growing and `len(T1)`
  lies. In a real cache that is a leak.
- **Checking T2 before T1 is fine; checking neither first is not.** Without the
  early return, the promoted key falls through and is re-added to T1.
- **Nil maps.** `New`-less construction panics on the first write; the test's
  literal initializes both.

## What the full algorithm adds

Real ARC also keeps ghost lists (B1, B2) of recently evicted keys and uses hits
on them to re-tune how much space T1 and T2 get. That adaptivity is the "A" in
ARC; the promotion rule here is its foundation.
