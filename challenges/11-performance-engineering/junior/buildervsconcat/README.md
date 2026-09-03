# Building A String Without The Garbage

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`s += part` in a loop allocates a fresh string on every pass and copies everything written so far — quadratic work and a heap full of dead intermediates. `strings.Builder` writes into one growing buffer, and if you tell it the final size it never grows at all.

## Task

Implement `JoinSep` in [buildervsconcat.go](buildervsconcat.go):

1. Concatenate `parts` with `sep` between adjacent elements.
2. Reach the result in a single allocation: compute the exact byte length first.
3. No parts yields `""`.

## Examples

**Example 1:**

```
Input:  JoinSep([a b], ", ")
Output: "a, b"
```

**Example 2:**

```
Input:  JoinSep([a], ", ")
Output: "a"
```

**Example 3:**

```
Input:  JoinSep(["" ""], "-")
Output: "-"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Strings are immutable** | Every `+=` produces a whole new string; the old one becomes garbage. |
| 2 | **`Builder.Grow`** | One sizing call turns a doubling buffer into a single exact allocation. |
| 3 | **`n-1` separators** | The size is the sum of the parts plus `len(sep)*(len(parts)-1)`. |

## Topics used again

`strings.Builder`, loops, integer arithmetic.

## Hint

Two passes: one to measure, one to write.

## Validate

```bash
make verify
```
