# Deque

## Intuition

The front insert is O(n) because every element shifts; that is the honest cost of a slice-backed deque, and it is fine while the deque stays small.

## Approach

1. `PushFront`: build a new slice starting with `v`.
2. `PushBack`: append.
3. `PopFront`: guard empty, read index 0, reslice.

## Solution

```go
func (d *Deque[T]) PushFront(v T) {
	d.items = append([]T{v}, d.items...)
}

func (d *Deque[T]) PushBack(v T) {
	d.items = append(d.items, v)
}

func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	front := d.items[0]
	d.items = d.items[1:]
	return front, true
}
```

## Walkthrough

`PushBack(1); PushFront(0)` produces `[0 1]`, so `PopFront` hands back `0`.

## Pitfalls

- Writing `append(d.items, v)` in `PushFront`, which pushes to the wrong end.
- Using a value receiver, so nothing is stored.
- Reslicing before reading the front element.
