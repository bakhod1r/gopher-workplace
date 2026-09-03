# Write Through The Caller's Array

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A helper meant to zero a scratch region builds a fresh slice of zeros and assigns it to its parameter. The scratch region is never actually zeroed.

## Task

Implement [fillshared.go](fillshared.go):

1. Set every element of `s` to `v`.
2. The writes must be visible through the caller's slice, including when `s` is a sub-slice view.
3. No allocations, no return value.

Replace the stub body in [fillshared.go](fillshared.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s := []int{1,2,3}; Fill(s, 7)
Output: s is [7 7 7]
```

**Example 2:**

```
Input:  s := []int{1,2,3,4}; Fill(s[1:3], 0)
Output: s is [1 0 0 4]
```

_Explanation:_ Only the view's range is written.

**Example 3:**

```
Input:  Fill(nil, 1)
Output: no panic
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slices share storage** | The parameter is a copy of the header, pointing at the same array. |
| 2 | **Views** | `s[1:3]` writes into the middle of the original array. |
| 3 | **Assignment vs mutation** | `s = ...` rebinds the local; `s[i] = ...` mutates the array. |

## Hint

Index into the parameter. Do not assign to it.

## Validate

```bash
make verify
```
