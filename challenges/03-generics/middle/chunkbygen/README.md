# Chunk By Boundary

**Level:** middle  
**Topic:** 03-generics

## Context

A log viewer groups consecutive lines belonging to the same request, splitting whenever the request ID changes.

## Task

Implement the stub(s) in [chunkbygen.go](chunkbygen.go):

1. Implement `ChunkBy`, starting a new group whenever `together(prev, cur)` is false.
2. Return an empty result for an empty input; never emit an empty group.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ChunkBy([]int{1,1,2}, equal)
Output: [][]int{{1,1},{2}}
```

**Example 2:**

```
Input:  ChunkBy([]int{1,2,3}, always)
Output: [][]int{{1,2,3}}
```

**Example 3:**

```
Input:  ChunkBy([]int{}, equal)
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Boundary predicates** | The predicate sees a pair, not a single element — it decides where the seam goes. |
| 2 | **Flushing the last group** | The final group is only appended after the loop ends. |
| 3 | **Stable order** | Appending in traversal order preserves the input's relative order. |

## Hint

Seed the first group with `s[0]`, then flush on each boundary and once more at the end.

## Validate

```bash
make verify
```
