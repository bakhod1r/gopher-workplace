# Generic Set

## Intuition

Unlike the slice-backed stack, a map-backed set must allocate before first write. That is why this type ships a constructor and the stack does not.

## Approach

1. `NewSet`: allocate the map inside the returned struct.
2. `Add`: assign `struct{}{}` at key `v`.
3. `Has`: use the comma-ok lookup.
4. `Len`: return `len(s.items)`.

## Solution

```go
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{items: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.items[v] = struct{}{}
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.items[v]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.items)
}
```

## Walkthrough

`Add(1); Add(1)` writes the same key twice, so `Len` stays `1` — the map enforces uniqueness.

## Pitfalls

- Returning `&Set[T]{}` without allocating the map, which panics on the first `Add`.
- Declaring `[T any]`, which cannot be a map key.
- Calling `NewSet()` without a type argument — nothing is there to infer from.
