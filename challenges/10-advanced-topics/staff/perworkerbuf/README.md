# Give Every Worker Its Own Frame

**Level:** staff
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A fan-out renderer was fixed for a data race by moving its scratch buffer inside the goroutine. The race is gone and the allocation count per row doubled — the fix reached for `make` when the frame would have done.

## Task

Implement [perworkerbuf.go](perworkerbuf.go):

1. Render every row concurrently as its values joined by `,`.
2. Return the results in input order.
3. Each goroutine's scratch must be a non-escaping local; no shared buffer, no per-row heap buffer.

Replace the stub body in [perworkerbuf.go](perworkerbuf.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  RenderAll([][]int{{1,2},{3}})
Output: ["1,2" "3"]
```

**Example 2:**

```
Input:  128 rows, 20 rounds
Output: every result correct
```

_Explanation:_ No goroutine touches another's scratch.

**Example 3:**

```
Input:  RenderAll(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Goroutine stacks** | Each goroutine has its own frame, so a local array is per-worker by construction. |
| 2 | **Escape analysis inside a goroutine** | A fixed-size array that only feeds `string(buf)` stays on the stack. |
| 3 | **Disjoint slot writes** | `out[i]` from goroutine i needs no synchronisation. |
| 4 | **Loop variables as parameters** | Passing `i` and `row` in keeps each goroutine's inputs its own. |

## Hint

A goroutine has a stack too. What lives there for free?

## Validate

```bash
make verify
```
