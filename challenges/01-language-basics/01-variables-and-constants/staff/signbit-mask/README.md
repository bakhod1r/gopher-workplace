# Sign Bit Mask

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

For an 8-bit two's-complement value the sign is bit 7 — mask `0x80`. The code
masks `0x40` (bit 6), so it misreads `0x40` and `0x80`.

## Task

Fix the single line between the markers in [sign.go](sign.go) so it tests the
correct sign bit.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Negative(0x80)
Output: true
```

_Explanation:_ Bit 7 set.

**Example 2:**

```
Input:  Negative(0x7F)
Output: false
```

**Example 3:**

```
Input:  Negative(0xFF)
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Sign bit position** | Bit 7 (`0x80`) for an 8-bit value. |
| 2 | **Bit mask** | `b & 0x80` isolates the top bit. |
| 3 | **Two's complement** | Sign bit set ⇒ negative. |

## Hint

Mask `0x80`.

## Validate

```bash
make verify
```
