# The Detour Through Interface Values

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A summing helper was made "generic" by routing the values through []any so one function could serve several element types. The generality was never used, and the allocation profile pays for it on every call.

## Task

Fix the single planted bug in [boxvalues.go](boxvalues.go):

1. Sum `vals` and return the total as an int64.
2. Fix the single bug so the function allocates nothing.
3. The accumulator must stay wide enough that large values do not overflow.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Total([]int{1,2,3})
Output: 6
```

**Example 2:**

```
Input:  Total(nil)
Output: 0
```

**Example 3:**

```
Input:  64 values
Output: 0 allocations, not 65
```

_Explanation:_ One for the slice, one per boxed element.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Interface boxing** | Storing an int in an `any` needs a heap word for the data half of the interface. |
| 2 | **The small-integer cache** | Values 0-255 reuse a runtime table, which is why the cost hides in toy tests. |
| 3 | **Escape through a container** | The `[]any` outlives each element's scope, so every box escapes. |
| 4 | **Generality has a price** | An unused abstraction still costs what it would have cost. |

## Hint

The second loop already has everything it needs. What is the first loop for?

## Validate

```bash
make verify
```
