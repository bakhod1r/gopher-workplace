# Byte Extraction Stride

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A byte is 8 bits, so byte `n` lives at bit offset `8*n`. Shifting by `4*n`
(a nibble stride) overlaps adjacent bytes — the truncation to `uint8` then hides
it for n=0 only.

## Task

Fix the single line between the markers in [extract.go](extract.go) so `ByteAt`
uses the correct 8-bit stride.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ByteAt(0x1122334455667788, 0)
Output: 0x88
```

**Example 2:**

```
Input:  ByteAt(0x1122334455667788, 1)
Output: 0x77
```

**Example 3:**

```
Input:  ByteAt(0x1122334455667788, 7)
Output: 0x11
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bit stride** | One byte = 8 bits, so offset is `8*n`. |
| 2 | **Conversion truncation** | `uint8(...)` keeps only the low 8 bits. |
| 3 | **Shift amount type** | Shift count is unsigned. |

## Hint

`v >> (8 * n)`.

## Validate

```bash
make verify
```
