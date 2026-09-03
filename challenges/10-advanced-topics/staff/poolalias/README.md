# The Result That Still Belongs To The Pool

**Level:** staff
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A wire encoder was pooled to cut allocations. Latency improved and, days later, a small fraction of responses started carrying another request's payload.

## Task

Fix the single planted bug in [poolalias.go](poolalias.go):

1. Render `vals` as decimal numbers joined by `,`.
2. The returned slice must be storage the caller owns.
3. The scratch buffer must still go back to the pool.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Encode([]int{1,2,3})
Output: "1,2,3"
```

**Example 2:**

```
Input:  a result held across 50 later calls
Output: unchanged
```

_Explanation:_ Otherwise the pool handed its bytes to someone else.

**Example 3:**

```
Input:  Encode(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ownership across a pool boundary** | `Put` transfers the buffer; anything still viewing it is now a dangling reference in spirit. |
| 2 | **Escape analysis and lifetime** | The result outlives the call; the scratch buffer must not. |
| 3 | **Silent corruption** | No panic, no race detector hit — just wrong bytes under concurrency. |

## Hint

Everything about the buffer is correct. Ask instead what the caller is holding when `Put` runs.

## Validate

```bash
make verify
```
