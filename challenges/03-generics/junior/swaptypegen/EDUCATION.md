# Swap Fields In Place

## Intuition

Mixing receiver kinds on one type is normal Go: pick the pointer only where mutation happens. The generic part changes nothing about that rule.

## Approach

1. `Swap`: assign the two fields to each other in one statement.
2. `Ordered`: return the smaller value first, leaving the fields alone.

## Solution

```go
func (p *SamePair[T]) Swap() {
	p.First, p.Second = p.Second, p.First
}

func (p SamePair[T]) Ordered() (T, T) {
	if p.Second < p.First {
		return p.Second, p.First
	}
	return p.First, p.Second
}
```

## Walkthrough

`SamePair[int]{2, 1}.Ordered()` returns `1, 2` while the stored pair still reads `{2, 1}`.

## Pitfalls

- Giving `Swap` a value receiver, so the exchange is lost.
- Making `Ordered` sort the fields in place, surprising the caller.
- Declaring `SamePair[A, B any]`, which would make the comparison impossible.
