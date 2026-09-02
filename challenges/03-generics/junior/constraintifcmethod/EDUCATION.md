# Constraint With A Method

## Intuition

Method constraints let generic code call behaviour without boxing. The caller keeps `[]Book`, and the compiler generates a direct call for that instantiation.

## Approach

1. Start a total at 0.
2. Add `it.Cents()` for each item.
3. Return the total.

## Solution

```go
func TotalCents[T Priced](items []T) int {
	total := 0
	for _, it := range items {
		total += it.Cents()
	}
	return total
}
```

## Walkthrough

`TotalCents([]Book{{200}, {350}})` instantiates `T = Book` and calls `Book.Cents` directly on each element.

## Pitfalls

- Taking `[]Priced` instead, forcing callers to convert their typed slice.
- Type-asserting inside the loop, which a constraint makes unnecessary.
- Returning a float — cents are integers for a reason.
