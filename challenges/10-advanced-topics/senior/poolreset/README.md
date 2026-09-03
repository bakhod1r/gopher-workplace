# The Pooled Buffer Nobody Emptied

**Level:** senior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A hot serialiser was converted to a `sync.Pool` and the throughput improved. Two days later the output of one request starts appearing at the front of another's.

## Task

Fix the single planted bug in [poolreset.go](poolreset.go):

1. Render `vals` as decimal numbers joined by `,`.
2. Fix the single bug so a borrowed buffer starts empty.
3. The buffer must still go back to the pool.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Render([]int{1,2,3})
Output: "1,2,3"
```

**Example 2:**

```
Input:  200 calls of Render([]int{7})
Output: every call returns "7"
```

_Explanation:_ A pooled buffer is not a fresh one.

**Example 3:**

```
Input:  the pooled buffer's capacity after 500 calls
Output: bounded
```

_Explanation:_ Otherwise every call appends to the last.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Pool semantics** | `Get` returns a value someone else put back, in whatever state they left it. |
| 2 | **Length vs capacity on reuse** | `[:0]` is what makes the capacity reusable and the contents gone. |
| 3 | **Cross-request contamination** | Reuse bugs leak one caller's data into another's output. |

## Hint

The first call is correct. The second one is not. What is different about the buffer it gets?

## Validate

```bash
make verify
```
