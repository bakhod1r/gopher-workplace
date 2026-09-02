# Strategy Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

The algorithm is the parameter. `Context` owns the data; the caller supplies the
per-element operation, so the same method doubles, squares or negates without
knowing which.

## Task

Implement `Process` on `*Context` in [strategypatt.go](strategypatt.go):

1. Apply `strategy` to every element of `c.Data`.
2. Write the results back **in place** — the caller inspects `c.Data`.

**Constraint (senior):** applying a strategy to 1M elements must happen in place, with zero allocations.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Data [1 2 3], strategy x*2
Output: Data [2 4 6]
```

**Example 2:**

```
Input:  Data [1 2 3], strategy x+1
Output: Data [2 3 4]
```

**Example 3:**

```
Input:  empty Data
Output: empty Data, strategy never called
```

_Explanation:_ nothing to transform.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Function parameter as strategy** | No interface required — `func(int) int` *is* the abstraction. |
| 2 | **`range` gives copies** | `for _, v := range` cannot write back; index the slice instead. |
| 3 | **In-place mutation** | Slices share their backing array, so writing through the index is visible to the caller. |

## Hint

`for i := range c.Data { c.Data[i] = strategy(c.Data[i]) }`. Assigning to the
`range` value variable changes nothing.

## Validate

```bash
make verify
```
