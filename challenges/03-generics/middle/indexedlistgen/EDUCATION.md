# List With Fast Membership

## Intuition

Two structures describing the same data is a classic trade: constant-time membership costs memory and the discipline of updating both together.

## Approach

1. `Append`: reject known values, record the position, then append.
2. `Has`: read the map.
3. `At`: bounds-check, then index the slice.

## Solution

```go
func NewIndexed[T comparable]() *Indexed[T] {
	return &Indexed[T]{items: make([]T, 0), index: make(map[T]int)}
}

func (l *Indexed[T]) Append(v T) bool {
	if _, ok := l.index[v]; ok {
		return false
	}
	l.index[v] = len(l.items)
	l.items = append(l.items, v)
	return true
}

func (l *Indexed[T]) Has(v T) bool {
	_, ok := l.index[v]
	return ok
}

func (l *Indexed[T]) At(i int) (T, bool) {
	if i < 0 || i >= len(l.items) {
		var zero T
		return zero, false
	}
	return l.items[i], true
}
```

## Walkthrough

The second `Append(a)` finds `a` in the index and returns `false` without touching the slice.

## Pitfalls

- Updating only one of the two structures.
- Scanning `items` in `Has`, which loses the whole benefit.
- Recording the position after appending, which is off by one.
