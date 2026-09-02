# Most Frequent

**Level:** middle  
**Topic:** 03-generics

## Context

An error report highlights the message that occurred most often, and must be reproducible across runs.

## Task

Implement the stub(s) in [modegen.go](modegen.go):

1. Implement `Mode`, returning the most frequent element and `true`.
2. On a tie return the element that appears earliest in `s`.
3. Return the zero value and `false` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Mode([]int{1,2,2})
Output: 2, true
```

**Example 2:**

```
Input:  Mode([]int{1,1,2,2})
Output: 1, true
```

**Example 3:**

```
Input:  Mode([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two passes** | Tally first, then pick — one pass cannot know the final counts. |
| 2 | **Deterministic ties** | Scanning `s` (not the map) for the winner avoids Go's randomised map order. |
| 3 | **Map iteration is unordered** | Reused from earlier: never break ties by ranging a map. |

## Hint

Break the tie by scanning the slice, not the map — map order is randomised.

## Validate

```bash
make verify
```
