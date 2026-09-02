# Top N

**Level:** middle  
**Topic:** 03-generics

## Context

A trending panel lists the five most common search terms, and the list must be identical between two runs over the same data.

## Task

Implement the stub(s) in [topngen.go](topngen.go):

1. Implement `TopN`, returning the `n` most frequent distinct values, most frequent first.
2. Break ties by first appearance in `s`.
3. Return an empty result for `n <= 0` or an empty input; a large `n` returns everything.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TopN([]int{1,2,2,3,3,3}, 2)
Output: [3 2]
```

**Example 2:**

```
Input:  TopN([]int{1,2}, 5)
Output: [1 2]
```

**Example 3:**

```
Input:  TopN([]int{1}, 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deterministic ordering** | First-appearance order plus a stable sort makes ties reproducible. |
| 2 | **Never sort the map** | Ranging a map for the ranking would randomise ties. |
| 3 | **Clamping `n`** | Asking for more than exists must not panic. |

## Hint

Collect the distinct values in first-seen order, then sort that slice stably by count.

## Validate

```bash
make verify
```
