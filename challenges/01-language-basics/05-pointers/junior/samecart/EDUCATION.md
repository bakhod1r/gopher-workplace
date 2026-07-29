# Struct identity via pointers

## Intuition

Pointer `==` compares addresses, distinguishing two structurally-equal instances.

## Approach

1. Compare the pointers with `==`.
2. True only when both name the same struct in memory.

## Solution

```go
type Cart struct{ Count int }

func Same(a, b *Cart) bool {
	return a == b
}
```

## Walkthrough

`Same(c, c)` compares one address to itself → `true`. Two `&Cart{}` literals allocate distinct structs → `false`.

## Pitfalls

- `*a == *b` compares fields; `a == b` compares identity.
- Distinct `&Cart{}` values are never equal pointers.
