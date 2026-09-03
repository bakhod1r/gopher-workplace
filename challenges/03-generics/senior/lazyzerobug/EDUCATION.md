# The Lazy Value Recomputed Forever

## Intuition

Using the zero value as the "not yet computed" marker is wrong for any `T` whose legitimate results include that zero: the cache never registers a hit.

## Approach

1. Check the `done` flag.
2. Compute and store when it is false.
3. Set `done` and return the stored value.

## Solution

```go
func (l *Lazy[T]) Get() T {
	if !l.done {
		l.value = l.Fn()
		l.done = true
	}
	return l.value
}
```

## Walkthrough

With `Fn` returning `0`, the comparison against the zero value is true on every call, so `Fn` runs again each time.

## Pitfalls

- Comparing through `any` at all — it panics for uncomparable types.
- Assuming callers never memoise a false, an empty string, or a nil pointer.
