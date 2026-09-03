# Give Back The Capacity You Stopped Using

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A long-lived index is built by appending millions of entries, then filtered down to a few thousand. The filtered index keeps the whole original array alive for the life of the process.

## Task

Implement [shrink.go](shrink.go):

1. Return a right-sized copy when `cap(s) > 2*len(s)`.
2. Return `s` untouched — and allocate nothing — otherwise.
3. A nil or empty input must not panic.

Replace the stub body in [shrink.go](shrink.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Shrink(make([]int,2,64))
Output: len 2, cap 2, new array
```

**Example 2:**

```
Input:  Shrink(make([]int,8,10))
Output: the same slice
```

_Explanation:_ 10 is not more than twice 8, so nothing is copied.

**Example 3:**

```
Input:  Shrink(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Capacity outlives length** | A short slice over a huge array pins the whole array. |
| 2 | **Right-sizing** | `make([]int, len(s))` plus `copy` releases the spare on the next collection. |
| 3 | **Copy only when it pays** | The threshold keeps the common case allocation-free. |

## Hint

`cap(s) <= 2*len(s)` is the cheap case. Return early.

## Validate

```bash
make verify
```
