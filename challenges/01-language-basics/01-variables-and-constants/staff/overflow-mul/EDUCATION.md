# Widen before you multiply

## Intuition

An arithmetic expression is evaluated in the operands' type, and the result
wraps **before** any surrounding conversion:

```go
var w, h int32 = 100000, 100000
int64(w * h) // w*h overflows int32 first, THEN widens the wrong value
int64(w) * int64(h) // widen first, multiply in int64 -> correct
```

The conversion `int64(...)` cannot recover bits already lost to the int32 wrap.

## Approach

1. `w * h` in `int32` overflows before widening.
2. Widen each factor first: `int64(w) * int64(h)`.

## Solution

```go
func Area(w, h int32) int64 {
	return int64(w) * int64(h)
}
```

## Walkthrough

`Area(100000, 100000)` overflows int32 in the bug; converting to int64 before multiplying keeps the full product.

## Pitfalls

- Convert **each operand** before the operation, not the result.
- Choose a result type wide enough for the maximum product.
- `go vet` will not catch this; it is valid code with wrong semantics.
