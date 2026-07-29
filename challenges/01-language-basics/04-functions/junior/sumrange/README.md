# Sum with For-Range

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`for range` walks a slice yielding index and value; ignore the index with `_`
when only the value matters.

## Task

Implement `SumRange` in [sumrange.go](sumrange.go) using a for-range loop.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumRange([1 2 3 4])
Output: 10
```

**Example 2:**

```
Input:  SumRange(nil)
Output: 0
```

**Example 3:**

```
Input:  SumRange([-1 1])
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **for-range** | `for _, v := range xs` iterates values. |
| 2 | **Blank identifier** | `_` discards the unused index. |
| 3 | **Accumulator** | A total starting at 0 grows each iteration. |

## Hint

`for _, v := range xs { total += v }`.

## Validate

```bash
make verify
```
