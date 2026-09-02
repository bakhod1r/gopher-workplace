# Generic Box

## Intuition

Checking `b.value == zero` would need `comparable` and would still be wrong: a deliberately stored zero is real data. An explicit flag keeps `T any` and stays correct.

## Approach

1. `Set`: assign the value and set `filled`.
2. `Get`: return the zero value and `false` while `filled` is false; otherwise the stored value and `true`.

## Solution

```go
func (b *Box[T]) Set(v T) {
	b.value = v
	b.filled = true
}

func (b *Box[T]) Get() (T, bool) {
	if !b.filled {
		var zero T
		return zero, false
	}
	return b.value, true
}
```

## Walkthrough

`Set(0); Get()` returns `0, true` because `filled` was set even though the value is the zero value.

## Pitfalls

- Detecting emptiness by comparing to the zero value.
- Using a value receiver on `Set`, which stores into a copy.
- Forgetting to set `filled`, so `Get` always reports `false`.
