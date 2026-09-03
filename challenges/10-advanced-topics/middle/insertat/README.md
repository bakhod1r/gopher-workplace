# Make Room In The Middle

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

An ordered list inserts by building a new slice from two sub-slices. Every insertion allocates and copies the whole list twice.

## Task

Implement [insertat.go](insertat.go):

1. Insert `v` at index `i`, shifting the rest up.
2. Clamp `i` into `[0, len(s)]`.
3. With spare capacity, allocate nothing.

Replace the stub body in [insertat.go](insertat.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  InsertAt([]int{1,3}, 1, 2)
Output: [1 2 3]
```

**Example 2:**

```
Input:  InsertAt([]int{1,2}, 99, 3)
Output: [1 2 3]
```

_Explanation:_ The index is clamped to the end.

**Example 3:**

```
Input:  InsertAt(nil, 0, 7)
Output: [7]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Grow first, then shift** | Appending a placeholder makes room for the copy. |
| 2 | **copy handles overlap** | A right shift over itself works because `copy` behaves like memmove. |
| 3 | **Capacity reuse** | The append is free when the room is already there. |

## Hint

Three steps: make room, shift, then write.

## Validate

```bash
make verify
```
