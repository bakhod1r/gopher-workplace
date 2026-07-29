# Checked Int64 Addition

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

An accumulator must detect int64 overflow instead of silently wrapping. The
positive-overflow test is `a > math.MaxInt64`, which is **never** true (nothing
exceeds the max), so `MaxInt64 + 1` wraps undetected.

## Task

Fix the positive-overflow check between the markers in
[checkedadd.go](checkedadd.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  math.MaxInt64, 1
Output: 0, false
```

_Explanation:_ Positive overflow detected.

**Example 2:**

```
Input:  math.MaxInt64-1, 1
Output: math.MaxInt64, true
```

**Example 3:**

```
Input:  math.MinInt64, -1
Output: 0, false
```

_Explanation:_ Negative overflow.

**Example 4:**

```
Input:  -5, -5
Output: -10, true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Overflow condition** | `a > MaxInt64 - b` when `b > 0`. |
| 2 | **Rearrange to avoid overflow** | Test with subtraction, not `a+b`. |
| 3 | **Symmetric negative case** | `a < MinInt64 - b` when `b < 0`. |

## Hint

`if b > 0 && a > math.MaxInt64-b`.

## Validate

```bash
make verify
```
