# Usable Zero Value

## Intuition

The read/write asymmetry of nil maps is exactly what makes lazy allocation cheap: one guard in one method buys a usable zero value.

## Approach

1. `Set`: allocate when nil, then assign.
2. `Get`: comma-ok lookup.
3. `Len`: return `len`.

## Solution

```go
func (s *Store[K, V]) Set(k K, v V) {
	if s.items == nil {
		s.items = make(map[K]V)
	}
	s.items[k] = v
}

func (s *Store[K, V]) Get(k K) (V, bool) {
	v, ok := s.items[k]
	return v, ok
}

func (s *Store[K, V]) Len() int {
	return len(s.items)
}
```

## Walkthrough

`var s Store[string,int]; s.Get("a")` reads a nil map and reports `false` without panicking.

## Pitfalls

- Guarding in `Get` too, which is dead code.
- Requiring a constructor and leaving the zero value broken.
- Allocating in `Len`, which turns a read into a mutation.
