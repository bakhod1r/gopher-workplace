# Versioned Value

## Intuition

Keeping the whole history rather than just the previous value makes repeated undo trivial, at the cost of memory that grows with every change.

## Approach

1. `Set`: append.
2. `Get`: read the last entry, or report `false`.
3. `Undo`: shrink by one when possible.
4. `Versions`: report the length.

## Solution

```go
func (v *Versioned[T]) Set(value T) {
	v.history = append(v.history, value)
}

func (v *Versioned[T]) Get() (T, bool) {
	if len(v.history) == 0 {
		var zero T
		return zero, false
	}
	return v.history[len(v.history)-1], true
}

func (v *Versioned[T]) Undo() bool {
	if len(v.history) == 0 {
		return false
	}
	v.history = v.history[:len(v.history)-1]
	return true
}

func (v *Versioned[T]) Versions() int {
	return len(v.history)
}
```

## Walkthrough

`Set(1); Set(2); Undo()` drops `2`, leaving `1` as the current value.

## Pitfalls

- Overwriting the last entry in `Set` instead of appending.
- Returning `true` from `Undo` on an empty history.
- Treating a stored zero as "never set".
