# Widen before you multiply

## The idea

An arithmetic expression is evaluated in the operands' type, and the result
wraps **before** any surrounding conversion:

```go
var w, h int32 = 100000, 100000
int64(w * h) // w*h overflows int32 first, THEN widens the wrong value
int64(w) * int64(h) // widen first, multiply in int64 -> correct
```

The conversion `int64(...)` cannot recover bits already lost to the int32 wrap.

## Why it matters

Products, areas, hashes, and accumulations routinely exceed the input type's
range. Signed overflow in Go wraps silently (no panic), producing a plausible
wrong number that passes small tests and fails in production at scale.

## Watch out

- Convert **each operand** before the operation, not the result.
- Choose a result type wide enough for the maximum product.
- `go vet` will not catch this; it is valid code with wrong semantics.

## Try it yourself

```go
var a, b int32 = 46341, 46341 // ~2^31
int64(a * b)         // wrong (wrapped)
int64(a) * int64(b)  // 2147488281
```
