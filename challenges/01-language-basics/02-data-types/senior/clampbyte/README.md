# Saturating Pixel Clamp

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

Brightening an image adds to pixel values; a result of `300` must saturate to
`255`, not wrap to `44`. `byte(x)` wraps — it keeps only the low 8 bits.

## Task

Fix the conversion between the markers in [clampbyte.go](clampbyte.go) to
saturate to [0,255].

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  128
Output: 128
```

_Explanation:_ in range, unchanged

**Example 2:**

```
Input:  300
Output: 255
```

_Explanation:_ saturates to max byte

**Example 3:**

```
Input:  -20
Output: 0
```

_Explanation:_ saturates to min byte

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Conversion wraps** | `byte(300) == 44` (mod 256). |
| 2 | **Saturation** | Clamp before converting. |
| 3 | **Range guard** | Two comparisons pin 0 and 255. |

## Hint

`if x < 0 {return 0}; if x > 255 {return 255}; return byte(x)`.

## Validate

```bash
make verify
```
