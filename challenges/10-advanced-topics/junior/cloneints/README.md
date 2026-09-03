# A Copy That Owns Its Memory

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A cache hands callers the slice it stores internally. Callers mutate what they get back, and the cached entry changes underneath the cache.

## Task

Implement [cloneints.go](cloneints.go):

1. Return a copy of `s` that shares no backing array with it.
2. Handle a nil or empty input without panicking.

Replace the stub body in [cloneints.go](cloneints.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Clone([]int{1,2,3})
Output: [1 2 3]
```

**Example 2:**

```
Input:  c := Clone(s); s[0] = 99
Output: c[0] is unchanged
```

_Explanation:_ The copy owns its own array.

**Example 3:**

```
Input:  Clone(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Slice header vs storage** | A slice is a pointer, length and capacity — assigning one copies the header, not the elements. |
| 2 | **copy** | `copy(dst, src)` moves elements up to the shorter length. |
| 3 | **Aliasing** | Two slices over one array see each other's writes. |

## Hint

`d := s` copies three words. What copies the elements?

## Validate

```bash
make verify
```
