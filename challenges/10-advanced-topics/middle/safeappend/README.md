# An Append That Cannot Reach The Caller's Tail

**Level:** middle
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

Two request handlers hold overlapping views of one buffer. One of them appends, and the other's data silently changes — the classic aliasing corruption that only shows under load.

## Task

Fix the single planted bug in [safeappend.go](safeappend.go):

1. Append `v` to `s` and return the result.
2. Elements of the backing array past `len(s)` must never be written.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Add([]int{1,2}, 3)
Output: [1 2 3]
```

**Example 2:**

```
Input:  b := []int{1,2,3,4}; Add(b[:2], 99)
Output: b is still [1 2 3 4]
```

_Explanation:_ The spare capacity belongs to someone else.

**Example 3:**

```
Input:  Add(nil, 1)
Output: [1]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Spare capacity is shared** | `append` writes into `cap` before it reallocates — and that memory may not be yours. |
| 2 | **Three-index slicing** | `s[:len(s):len(s)]` sets cap == len, so the next append must allocate. |
| 3 | **Slice headers** | Two headers over one array is normal; only capacity makes it dangerous. |

## Hint

`append` reallocates only when the capacity runs out. Make it run out.

## Validate

```bash
make verify
```
