# One Scratch Buffer, Many Goroutines

**Level:** staff
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A fan-out encoder passes its unit tests, passes review, and produces interleaved garbage the first time it runs with more than one core busy.

## Task

Fix the single planted bug in [sharedscratch.go](sharedscratch.go):

1. Render every batch concurrently as decimal numbers joined by `,`.
2. Return the results in input order.
3. Fix the single bug so each goroutine writes only to memory it owns.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  EncodeAll([][]int{{1,2},{3}})
Output: ["1,2" "3"]
```

**Example 2:**

```
Input:  64 batches, 20 rounds
Output: every result correct every time
```

_Explanation:_ No goroutine may touch another's scratch.

**Example 3:**

```
Input:  EncodeAll(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Captured variables are shared** | A closure over an outer variable gives every goroutine the same variable. |
| 2 | **Data race vs logic race** | Two goroutines appending to one slice corrupt both results even when no write tears. |
| 3 | **Disjoint writes are safe** | Writing `out[i]` from goroutine i needs no lock — the elements do not overlap. |

## Hint

`i` and `b` were passed in as parameters for a reason. What else does the closure use?

## Validate

```bash
make verify
```
