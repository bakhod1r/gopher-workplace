# No Field Access On T

## Intuition

Go deliberately has no structural field constraint, which is why every library helper of this shape takes a `func(T) K` instead.

## Approach

1. Sum `price(v)` across the items.
2. Return the total.

## Solution

```go
func TotalPrice[T any](items []T, price func(T) int) int {
	total := 0
	for _, v := range items {
		total += price(v)
	}
	return total
}
```

## Walkthrough

`TotalPrice(books, bookPrice)` works for any element type because the field access happens in the caller's function, where the type is concrete.

## Pitfalls

- Trying `v.Price` on a type parameter.
- Inventing a constraint listing struct fields — no such thing exists.
- Falling back to reflection when a projection is simpler and faster.
