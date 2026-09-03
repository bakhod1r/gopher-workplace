# Chunks That Share Storage

**Level:** senior  
**Topic:** 03-generics

## Context

A batching layer hands each chunk to a worker that appends a trailer before sending. In production the first record of the following batch keeps getting clobbered.

## Task

Fix the single planted bug in [aliasbug.go](aliasbug.go):

1. Find and fix the single bug so each returned group is independent of the input.
2. Appending to one group must not change the input or any other group.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Chunk([]int{1,2,3,4}, 2)
Output: [][]int{{1,2},{3,4}}
```

**Example 2:**

```
Input:  append into group 0
Output: input unchanged
```

**Example 3:**

```
Input:  Chunk([]int{1}, 0)
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backing-array aliasing** | A sub-slice shares storage with its source; `append` into it can overwrite neighbours. |
| 2 | **Capacity leaks past the length** | `s[i:end]` keeps the source capacity, so `append` writes into the next group. |
| 3 | **Copy or clip** | Either copy the window or cap the capacity with a three-index slice. |

## Hint

The groups look right until someone appends to one of them.

## Validate

```bash
make verify
```
