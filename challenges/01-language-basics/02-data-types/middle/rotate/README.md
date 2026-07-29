# Circular Bit Rotation

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A rotate is a shift where the bits that fall off one end wrap to the other. For
a byte, rotating by `n` is really rotating by `n % 8`.

## Task

Implement `Left(b, n)` rotating the 8 bits of `b` left by `n` (n may be ≥ 8).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Left(0b1000_0000, 1)
Output: 0b0000_0001
```

_Explanation:_ top bit wraps to bottom

**Example 2:**

```
Input:  Left(0b1010_0000, 4)
Output: 0b0000_1010
```

_Explanation:_ nibbles swapped

**Example 3:**

```
Input:  Left(0b0000_0001, 9)
Output: 0b0000_0010
```

_Explanation:_ n masked to 9&7=1

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shift + OR** | `(b<<n) |
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
