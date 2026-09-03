# Batch A Slice Into Fixed-Size Windows

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A batch uploader sends records in groups of 500. The current splitter copies every record into a fresh group and doubles the job's peak memory.

## Task

Implement [chunk.go](chunk.go):

1. Split `s` into consecutive groups of at most `n`.
2. The groups must be views into `s`, not copies.
3. Return nil when `n <= 0`; preallocate the outer slice to its exact group count.

Replace the stub body in [chunk.go](chunk.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Chunk([]int{1,2,3,4,5}, 2)
Output: [[1 2] [3 4] [5]]
```

_Explanation:_ The last group holds the remainder.

**Example 2:**

```
Input:  Chunk([]int{1,2}, 5)
Output: [[1 2]]
```

_Explanation:_ n larger than the input gives one group.

**Example 3:**

```
Input:  Chunk([]int{1}, 0)
Output: <nil>
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sub-slicing** | `s[i:end]` is a view — no elements move. |
| 2 | **Ceiling division** | `(len+n-1)/n` is the group count, so the outer slice is sized once. |
| 3 | **Boundary clamping** | The final window must stop at `len(s)`. |

## Hint

Two allocations at most: the outer slice, and nothing else.

## Validate

```bash
make verify
```
