# Absolute Value

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A scoreboard computes point differences. Negative differences should be shown
as positive distances.

## Task

Implement `Abs` on `MyInt` in [absval.go](absval.go):

1. Return the absolute value of the integer.
2. `Abs` returns `MyInt`, not `int`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MyInt(-5).Abs()
Output: 5
```

**Example 2:**

```
Input:  MyInt(3).Abs()
Output: 3
```

**Example 3:**

```
Input:  MyInt(0).Abs()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods on defined types** | `type MyInt int` has its own method set. |
| 2 | **Value receiver** | Read-only — returns a new value. |
| 3 | **Conditional logic** | Simple `if` for negative check. |

## Hint

`if n < 0 { return -n }; return n` — negate if negative.

## Validate

```bash
make verify
```
