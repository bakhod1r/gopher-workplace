# Circular Bit Rotation

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A rotate is a shift where the bits that fall off one end wrap to the other. For
a byte, rotating by `n` is really rotating by `n % 8`.

## Task

Implement `Left(b, n)` rotating the 8 bits of `b` left by `n` (n may be ≥ 8).

## Examples

```go
Left(0b1000_0000, 1) // => 0b0000_0001
Left(0b1010_0000, 4) // => 0b0000_1010
Left(0x01, 9)        // => 0x02
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shift + OR** | `(b<<n) | (b>>(8-n))` merges the wrapped bits. |
| 2 | **Modulo width** | Reduce `n` mod 8 first. |
| 3 | **byte width** | Mask/rely on `byte` truncation to 8 bits. |

## Hint

`n &= 7; return b<<n | b>>(8-n)` — and handle `n==0` (shift by 8 is undefined
intent; the mask makes it 0, so guard or the expression still works since
`b>>8==0`).

## Validate

```bash
make verify
```
