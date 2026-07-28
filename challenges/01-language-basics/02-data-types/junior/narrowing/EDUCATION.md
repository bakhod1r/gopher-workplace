# Narrowing conversions

## The idea

Converting to a smaller integer type keeps only the low bits — it **wraps**, it
does not clamp or error:

```go
int8(200)  // -56 (200 mod 256, reinterpreted as signed)
uint8(300) // 44
```

## Why it matters

Narrowing happens whenever you pack a wider value into a smaller field (bytes,
protocol fields, pixel channels). If the value is out of range, you get a silent
wrong number, not a panic — so validate before narrowing.

## Watch out

- Conversions never panic on overflow; they truncate bits.
- Signedness matters: `int8(200)` reinterprets the high bit as sign.
- Check the range first, or clamp, when the input may exceed the target type.
