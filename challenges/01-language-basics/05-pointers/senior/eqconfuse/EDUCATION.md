# Pointer identity vs pointed-to equality

## Intuition

`a == b` tests whether two pointers share an address; `*a == *b` tests whether their values match — different questions.

## Approach

1. Identity means same **address**, not equal values.
2. The bug compares `*a == *b` (values). Compare the pointers: `a == b`.

## Solution

```go
func Same(a, b *int) bool {
	return a == b
}
```

## Walkthrough

`Same(&x, &y)` with equal values returns true under the bug (`*a == *b`), but they are distinct variables. Comparing `a == b` reports identity correctly.

## Pitfalls

- `*a == *b` is true for two different vars holding 5.
- Use `a == b` for identity.
