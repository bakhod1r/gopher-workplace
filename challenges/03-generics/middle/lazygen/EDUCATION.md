# Lazy Value

## Intuition

The flag keeps the type usable for every `T`, including types that are not comparable at all — a zero-value check would need `comparable`.

## Approach

1. `Get`: compute and record when not done, then return the stored value.
2. `Done`: report the flag.

## Solution

```go
func NewLazy[T any](compute func() T) *Lazy[T] {
	return &Lazy[T]{compute: compute}
}

func (l *Lazy[T]) Get() T {
	if !l.done {
		l.value = l.compute()
		l.done = true
	}
	return l.value
}

func (l *Lazy[T]) Done() bool {
	return l.done
}
```

## Walkthrough

A `compute` returning `0` still flips `done`, so the second `Get` reuses the stored zero rather than recomputing.

## Pitfalls

- Comparing the stored value to the zero value to decide whether to compute.
- Recomputing on every call and defeating the point.
- Claiming concurrency safety without `sync.Once`.
