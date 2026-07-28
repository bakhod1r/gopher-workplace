# abs() and the most-negative value

## The idea

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

## Why it matters

Naive `abs` (`if x < 0 { return -x }`) is wrong for `MinInt`. This bites sort
comparators, distance metrics, and hashing. The fix is about *when* you widen,
not the branch logic.

## Watch out

- `-int8(-128)` overflows; `-int(int8(-128))` does not.
- Standard libraries often document that `abs(MinInt)` is undefined/overflowing —
  handle it explicitly if it can occur.
- The same asymmetry breaks `x / -1` for MinInt in some languages.

## Try it yourself

```go
var m int8 = -128
if m < 0 { fmt.Println(-int(m)) } // 128, correct
if m < 0 { fmt.Println(int(-m)) } // -128, wrong
```
