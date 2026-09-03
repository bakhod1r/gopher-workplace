# Did That Append Reallocate

**Level:** junior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A reviewer claims a hot loop never reallocates because it appends to a pre-sized slice. Nobody has actually measured it.

## Task

Implement [capgrew.go](capgrew.go):

1. Append `v` to `s` and return the result.
2. Report whether the capacity changed — that is, whether a new array was allocated.

Replace the stub body in [capgrew.go](capgrew.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Appended(make([]int,0,4), 1)
Output: [1], false
```

_Explanation:_ The room was already there.

**Example 2:**

```
Input:  Appended(make([]int,1,1), 2)
Output: [.. 2], true
```

_Explanation:_ The capacity was exhausted.

**Example 3:**

```
Input:  Appended(nil, 5)
Output: [5], true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Capacity is the growth signal** | `append` reallocates exactly when the length would exceed the capacity. |
| 2 | **Length always changes** | Only the capacity distinguishes a copy from an in-place write. |
| 3 | **Reading cap before and after** | The comparison is the whole measurement. |

## Hint

Record something before the append and compare it afterwards.

## Validate

```bash
make verify
```
