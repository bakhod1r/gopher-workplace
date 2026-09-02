# Generic Counter

## Intuition

Map reads of missing keys yield the zero value, so incrementing needs no existence check — the same reason `Count` can return the raw lookup.

## Approach

1. `NewCounter`: allocate the map.
2. `Inc`: `c.counts[v]++`.
3. `Count`: return the map lookup.
4. `Total`: sum the values.

## Solution

```go
func NewCounter[T comparable]() *Counter[T] {
	return &Counter[T]{counts: make(map[T]int)}
}

func (c *Counter[T]) Inc(v T) {
	c.counts[v]++
}

func (c *Counter[T]) Count(v T) int {
	return c.counts[v]
}

func (c *Counter[T]) Total() int {
	total := 0
	for _, n := range c.counts {
		total += n
	}
	return total
}
```

## Walkthrough

`Inc("a")` on a fresh counter reads `0`, adds one, and stores `1`.

## Pitfalls

- Guarding `Inc` with a `_, ok :=` check that changes nothing.
- Returning `len(c.counts)` from `Total`, which counts distinct keys instead of events.
- Forgetting the constructor and panicking on a nil map write.
