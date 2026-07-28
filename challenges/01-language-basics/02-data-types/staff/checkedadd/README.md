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

## Examples

```go
Add(math.MaxInt64, 1) // => 0, false
Add(math.MaxInt64-1,1)// => MaxInt64, true
```

## Topics to Master

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
