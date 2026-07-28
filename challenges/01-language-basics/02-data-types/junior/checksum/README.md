# Checksum

**Level:** junior
**Topic:** 01-language-basics → 02-data-types
**Estimated time:** 15 min

## Context

A wire protocol appends a one-byte checksum to each frame: the sum of the payload
bytes, kept in a single `uint8`. Because the field is exactly 8 bits wide, the
sum naturally wraps around past 255 — that overflow is part of the spec, not a
bug.

## Task

Implement `Checksum` in [checksum.go](checksum.go) so that it returns the
modulo-256 sum of the bytes in `data`, accumulated in a `uint8`:

1. Add every byte together; the running total is a `uint8`.
2. When the total exceeds 255 it wraps around (`255 + 1 == 0`).
3. A nil or empty slice returns `0`.

Do **not** change the function signature or the tests.

## Examples

```go
Checksum([]byte{1, 2, 3})  // => 6
Checksum([]byte{255, 1})   // => 0   (256 wraps to 0)
Checksum([]byte{200, 100}) // => 44  (300 mod 256)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Fixed-width integers** | `uint8` holds 0–255; arithmetic that exceeds the width wraps modulo 2⁸ instead of growing. |
| 2 | **`byte` is `uint8`** | Elements of a `[]byte` are already `uint8`, so summing them in a `uint8` accumulator wraps automatically. |
| 3 | **Zero value** | A freshly declared `uint8` accumulator starts at `0` — the right identity for a sum. |

## Hint

Keep the accumulator itself a `uint8`. If you sum into an `int` and only convert
at the end, you will get the wrong answer unless you also take `% 256`. Summing
directly in a `uint8` gives the wrap-around for free.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
