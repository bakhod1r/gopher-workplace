# Chunk

**Level:** junior  
**Topic:** 03-generics

## Context

A batch uploader sends at most N records per request. The final batch is usually short.

## Task

Implement the stub(s) in [chunkgen.go](chunkgen.go):

1. Implement `Chunk`, splitting `s` into consecutive groups of at most `size` elements.
2. Return an empty result when `size <= 0` or `s` is empty.
3. Each group must be an independent slice, not a window into `s`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Chunk([]int{1, 2, 3}, 2)
Output: [][]int{{1, 2}, {3}}
```

**Example 2:**

```
Input:  Chunk([]int{1, 2}, 5)
Output: [][]int{{1, 2}}
```

**Example 3:**

```
Input:  Chunk([]int{1}, 0)
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slices of slices** | `[][]T` is a slice whose element type is itself a slice of the type parameter. |
| 2 | **Sharing a backing array** | Reused from language basics: `s[i:end]` aliases `s`; `copy` breaks the link. |
| 3 | **Clamping the tail** | The last group is short whenever `len(s)` is not a multiple of `size`. |

## Hint

Step `i` by `size` and clamp `end` to `len(s)`.

## Validate

```bash
make verify
```
