# Widen Before Multiply

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`int64(w * h)` computes `w*h` in `int32` first — it overflows and wraps, and the
widening conversion happens *after* the damage. Widen the operands first.

## Task

Fix the single line between the markers in [area.go](area.go) so large products
do not overflow.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Area(100000, 100000)
Output: 10000000000
```

**Example 2:**

```
Input:  Area(2, 3)
Output: 6
```

**Example 3:**

```
Input:  Area(46341, 46341)
Output: no 32-bit overflow
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Operand type wins** | `w*h` is `int32`; the result wraps before conversion. |
| 2 | **Widen first** | Convert each operand to `int64` before `*`. |
| 3 | **Silent overflow** | Signed wraparound gives a plausible wrong number. |

## Hint

`int64(w) * int64(h)`.

## Validate

```bash
make verify
```
