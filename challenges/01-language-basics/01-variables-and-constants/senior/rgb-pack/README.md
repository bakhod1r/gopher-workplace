# RGB Bit Packing

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Packing RGB into `0x00RRGGBB` needs red shifted 16 bits, green 8, blue 0. The
code shifts red by only 8, colliding it with green.

## Task

Fix the single line between the markers in [color.go](color.go) so red occupies
bits 16–23.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Pack(0x12, 0x34, 0x56)
Output: 0x123456
```

**Example 2:**

```
Input:  Red(0x123456)
Output: 0x12
```

**Example 3:**

```
Input:  Pack(255, 0, 0)
Output: 0xFF0000
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bit shifting** | Each channel occupies its own 8-bit lane. |
| 2 | **Conversion width** | Convert `uint8` to `uint32` before shifting past 8 bits. |
| 3 | **OR composition** | ` |

## Hint

Red must shift 16: `uint32(r)<<16 | uint32(g)<<8 | uint32(b)`.

## Validate

```bash
make verify
```
