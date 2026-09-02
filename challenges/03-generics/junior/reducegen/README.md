# Reduce

**Level:** junior  
**Topic:** 03-generics

## Context

Totals, joins, and maximums all share one shape: carry an accumulator across the slice and combine it with each element.

## Task

Implement the stub(s) in [reducegen.go](reducegen.go):

1. Implement `Reduce`, starting from `init` and applying `f(acc, e)` for each element left to right.
2. Return `init` unchanged for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Reduce([]int{1, 2, 3}, 0, add)
Output: 6
```

**Example 2:**

```
Input:  Reduce([]string{"a", "b"}, "", concat)
Output: "ab"
```

**Example 3:**

```
Input:  Reduce([]int{}, 5, add)
Output: 5
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Accumulator type parameter** | `A` is independent of `T`, so you can fold `[]int` into a `string`. |
| 2 | **Functions as values** | Reused from language basics: a `func(T) U` parameter is an ordinary value. |
| 3 | **Left fold order** | `f` is applied in index order — that matters for non-commutative operations like string concatenation. |

## Hint

One local variable, reassigned each iteration.

## Validate

```bash
make verify
```
