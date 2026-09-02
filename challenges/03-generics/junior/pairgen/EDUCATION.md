# Generic Pair

## Intuition

A method may name any instantiation of its own generic type, including one with the parameters reordered. What it may not do is introduce a brand-new type parameter of its own.

## Approach

1. `MakePair`: return a composite literal with both fields set.
2. `Swapped`: return `Pair[B, A]` with the values crossed over.

## Solution

```go
func (p Pair[A, B]) Swapped() Pair[B, A] {
	return Pair[B, A]{First: p.Second, Second: p.First}
}

func MakePair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{First: a, Second: b}
}
```

## Walkthrough

`MakePair(1, "a").Swapped()` produces `Pair[string, int]{"a", 1}` — the compiler tracks the type swap for you.

## Pitfalls

- Declaring `func (p Pair[A, B]) Swapped[C any]()` — methods cannot add type parameters.
- Returning `Pair[A, B]` with the values swapped, which does not type-check.
- Mutating the receiver instead of returning a new pair.
