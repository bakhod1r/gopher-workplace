# Where The Padding Comes From

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Struct layout is not a mystery, it is two rules applied in order: every field starts at a multiple of its alignment, and the struct's total size is rounded up to its widest field's alignment. Implement those rules once and you can predict any struct's size without running the program.

## Task

Implement both functions in [fieldpad.go](fieldpad.go):

1. `AlignUp` rounds `n` up to the next multiple of `a`; an `a` of `0` or `1` leaves `n` alone.
2. `StructSize` lays fields out in order, aligning each to its own size.
3. `StructSize` pads the total to the largest field's alignment; zero-sized fields are skipped and no fields gives `0`.

## Examples

**Example 1:**

```
Input:  AlignUp(9, 8)
Output: 16
```

**Example 2:**

```
Input:  StructSize([8 4 2 1])
Output: 16
```

**Example 3:**

```
Input:  StructSize([1 8 2 4])
Output: 24
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Alignment rounding** | `(n + a - 1) / a * a` is the whole rule. |
| 2 | **Trailing padding** | The struct is padded so an array of it keeps every element aligned. |
| 3 | **Order changes the answer** | The same field sizes give 16 or 24 depending only on declaration order. |

## Topics used again

Loops, integer arithmetic, `max`.

## Hint

Track the running offset and the largest field seen; align at each step and once more at the end.

## Validate

```bash
make verify
```
