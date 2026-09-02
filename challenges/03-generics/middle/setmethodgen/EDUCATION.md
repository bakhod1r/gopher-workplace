# Set With Operations

## Intuition

Chained set algebra only reads correctly if each step is pure; mutating the receiver would make `a.Union(b)` change `a` and break every later use.

## Approach

1. `NewSet`: allocate and fill from the variadic arguments.
2. `Union`: copy both sides into a new set.
3. `Intersect`: copy only the elements present in both.
4. `Len`: report the size.

## Solution

```go
func NewSet[T comparable](vs ...T) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{}, len(vs))}
	for _, v := range vs {
		s.items[v] = struct{}{}
	}
	return s
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	for k := range s.items {
		out.items[k] = struct{}{}
	}
	for k := range other.items {
		out.items[k] = struct{}{}
	}
	return out
}

func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	out := NewSet[T]()
	for k := range s.items {
		if _, ok := other.items[k]; ok {
			out.items[k] = struct{}{}
		}
	}
	return out
}

func (s *Set[T]) Len() int {
	return len(s.items)
}
```

## Walkthrough

`NewSet(1,2).Union(NewSet(2,3))` holds `{1,2,3}`, and both operands still hold what they did.

## Pitfalls

- Writing results into the receiver, corrupting it for later expressions.
- Returning `Set[T]` by value and copying the map header around.
- Assuming the receiver is non-nil without documenting it.
