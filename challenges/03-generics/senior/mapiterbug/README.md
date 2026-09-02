# A Ranking That Changes Between Runs

**Level:** senior  
**Topic:** 03-generics

## Context

An error report highlights a different "top error" on every run, even for identical input, and nobody can reproduce the screenshots.

## Task

Fix the single planted bug in [mapiterbug.go](mapiterbug.go):

1. Find and fix the single bug so ties resolve to the earliest element in `s`.
2. The result must be identical across runs for identical input.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Mode([]int{1,1,2,2})
Output: 1, true
```

**Example 2:**

```
Input:  Mode([]int{1,2,2})
Output: 2, true
```

**Example 3:**

```
Input:  Mode([]int{})
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deterministic output** | Go randomises map range order; ranking must come from a slice. |
| 2 | **Randomisation is deliberate** | Go shuffles map iteration so code cannot depend on an accidental order. |
| 3 | **Tally then rank** | Count in the map, but decide the winner by walking the input. |

## Hint

Where does the loop that picks the winner get its order from?

## Validate

```bash
make verify
```
