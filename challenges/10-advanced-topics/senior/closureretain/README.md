# The Callback That Held The Whole Batch

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A pipeline stage returns a small "report this later" callback per batch. The callbacks are tiny, there are thousands of them, and the heap holds every batch that ever passed through.

## Task

Fix the single planted bug in [closureretain.go](closureretain.go):

1. Return a function reporting the batch's total `Size`.
2. The total must be fixed when `Summarize` returns — later edits to the batch are invisible.
3. Fix the single bug so the callback does not keep the batch alive.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  f := Summarize(batch); f()
Output: the total
```

**Example 2:**

```
Input:  batch[0].Size = 100 after Summarize
Output: f() unchanged
```

_Explanation:_ The total is a snapshot.

**Example 3:**

```
Input:  heap with the callback alive
Output: the batch is collectable
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closures capture variables, not values** | Capturing the slice keeps its backing array reachable. |
| 2 | **Retention through a callback** | A one-word result held by a closure over megabytes still costs megabytes. |
| 3 | **Compute eagerly to release early** | Capturing the answer breaks the reference to the data. |

## Hint

The loop is in the wrong place. Where does the answer have to be computed for the batch to be droppable?

## Validate

```bash
make verify
```
