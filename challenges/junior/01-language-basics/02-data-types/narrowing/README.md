# Narrowing Conversion

**Level:** junior
**Topic:** 01-language-basics → 02-data-types
**Estimated time:** 10 min

## Context

A service reads 64-bit counters but a legacy protocol field is only 32-bit
wide. The team needs a conversion that matches exactly what the wire format
does when a value overflows — Go's numeric narrowing, not a saturating clamp.

## Task

Implement `ToInt32` in [narrowing.go](narrowing.go) so it converts an `int64`
to an `int32`, wrapping around (two's-complement) when the value does not fit.

Do **not** change the function signature or the tests.

## Examples

```go
ToInt32(42)         // => 42
ToInt32(-7)         // => -7
ToInt32(2147483647) // => 2147483647   (int32 max)
ToInt32(2147483648) // => -2147483648  (wraps to int32 min)
ToInt32(4294967296) // => 0            (2^32 wraps to 0)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Explicit conversion** | Go never converts between numeric types implicitly; you write `int32(n)`. |
| 2 | **Narrowing** | Converting to a smaller type keeps only the low bits — the high bits are discarded. |
| 3 | **Two's-complement wrap** | Dropping the high bits makes an out-of-range value wrap, not clamp: one past `int32` max becomes `int32` min. |

## Hint

The whole task is a single explicit conversion, `int32(n)`. Go's spec already
defines the wrap-around behaviour — you do not add any range checks or clamping.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
