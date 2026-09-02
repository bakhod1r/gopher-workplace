# The Named Type That Loses Its Width

**Level:** staff  
**Topic:** 03-generics

## Context

A latency roll-up stores durations as a named `Millis` type over `int64`. Totals are correct in staging and wildly negative in production, where the numbers are larger.

## Task

Fix the single planted bug in [tildeconvbug.go](tildeconvbug.go):

1. Find and fix the single bug so the sum is accumulated at the width the constraint promises.
2. Every named type with underlying type `int64` must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Total([]Millis{1, 2, 3})
Output: 6
```

**Example 2:**

```
Input:  Total([]Millis{3000000000, 3000000000})
Output: 6000000000
```

**Example 3:**

```
Input:  Total([]Millis{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Underlying types** | `~int64` admits every named type whose underlying type is `int64` — the values are still full width. |
| 2 | **Scale is a requirement** | A graded test asserts the result on millions of elements, so a defect that only shows past a threshold is caught. |
| 3 | **Conversions truncate silently** | `int32(v)` discards the high bits with no error and no panic. |

## Hint

What width is the accumulator?

## Validate

```bash
make verify
```
