# Methods Take No Type Parameters

## Intuition

Method sets have to be fully known when a type is instantiated, so per-method type parameters cannot exist — this is why the stdlib exposes `slices.Map`-shaped helpers as functions.

## Approach

1. `MapBag`: build a `Bag[U]` from the converted items.
2. `Add`: copy, append, and return a new bag.
3. `Items`: return a copy.

## Solution

```go
func MapBag[T, U any](b Bag[T], f func(T) U) Bag[U] {
	out := Bag[U]{items: make([]U, 0, len(b.items))}
	for _, v := range b.items {
		out.items = append(out.items, f(v))
	}
	return out
}

func (b Bag[T]) Add(v T) Bag[T] {
	items := make([]T, len(b.items), len(b.items)+1)
	copy(items, b.items)
	return Bag[T]{items: append(items, v)}
}

func (b Bag[T]) Items() []T {
	out := make([]T, len(b.items))
	copy(out, b.items)
	return out
}
```

## Walkthrough

`MapBag(bag, itoa)` returns a `Bag[string]` — an instantiation no method on `Bag[int]` could have named.

## Pitfalls

- Trying `func (b Bag[T]) Map[U any](...)`, which does not compile.
- Mutating the receiver's slice in `Add`, which surprises value-semantics callers.
- Returning the internal slice from `Items`.
