# Generic List

## Intuition

A negative index panics just as a too-large one does, and it is the case people forget. Checking both makes `At` safe for any caller input.

## Approach

1. `Append`: append to the slice.
2. `At`: return `zero, false` when out of range, otherwise the element.
3. `Len`: return the slice length.

## Solution

```go
func (l *List[T]) Append(v T) {
	l.items = append(l.items, v)
}

func (l *List[T]) At(i int) (T, bool) {
	if i < 0 || i >= len(l.items) {
		var zero T
		return zero, false
	}
	return l.items[i], true
}

func (l *List[T]) Len() int {
	return len(l.items)
}
```

## Walkthrough

`At(-1)` fails the first half of the guard and returns `0, false` rather than panicking.

## Pitfalls

- Checking only the upper bound.
- Returning `items[i]` before the guard.
- Using a value receiver on `Append`.
