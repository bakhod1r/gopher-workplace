# abs() and the most-negative value

## Intuition

Because of two's-complement asymmetry, the most-negative value has no positive
counterpart in the same width. Negating it overflows and returns itself:

```go
var x int8 = -128
-x        // still -128 (128 doesn't fit int8)
int(-x)   // -128 — the overflow happened before widening
```

Widen **first**, then negate in the larger type where the result fits:

```go
-int(x)   // 128
```

## Approach

1. Negating `int8(-128)` overflows int8.
2. Widen first: `-int(x)` computes in `int`, not `int8`.

## Solution

```go
func Abs(x int8) int {
	if x < 0 {
		return -int(x)
	}
	return int(x)
}
```

## Walkthrough

`-x` on -128 stays -128 (int8 overflow); `-int(x)` widens to int then negates → 128.

## Pitfalls

- `-int8(-128)` overflows; `-int(int8(-128))` does not.
- Standard libraries often document that `abs(MinInt)` is undefined/overflowing —
  handle it explicitly if it can occur.
- The same asymmetry breaks `x / -1` for MinInt in some languages.
