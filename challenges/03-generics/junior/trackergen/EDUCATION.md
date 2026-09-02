# Min Max Tracker

## Intuition

Streaming means there is no `s[0]` to seed from up front, so the count field carries that job — and it doubles as the emptiness flag for `Bounds`.

## Approach

1. `Add`: on the first value set `lo`, `hi`, and the count; afterwards compare and update.
2. `Bounds`: report `false` while the count is zero.

## Solution

```go
func (t *Tracker[T]) Add(v T) {
	if t.n == 0 {
		t.lo, t.hi = v, v
		t.n = 1
		return
	}
	if v < t.lo {
		t.lo = v
	}
	if v > t.hi {
		t.hi = v
	}
	t.n++
}

func (t *Tracker[T]) Bounds() (T, T, bool) {
	if t.n == 0 {
		var zero T
		return zero, zero, false
	}
	return t.lo, t.hi, true
}
```

## Walkthrough

`Add(3); Add(1)` seeds `lo = hi = 3`, then lowers `lo` to `1`, leaving `hi` at `3`.

## Pitfalls

- Comparing the first value against zero-valued bounds, which breaks for all-positive data.
- Using a value receiver, so nothing is ever recorded.
- Reporting `true` from `Bounds` before any value was added.
