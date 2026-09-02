# Find Value

**Level:** junior
**Topic:** 04-error-handling

## Context

A routing table maps a numeric code to a handler slot. An unknown code is a lookup failure, not slot zero.

## Task

Implement `Find` in [mustfind.go](mustfind.go):

1. Return the index of the first occurrence of `target` and nil.
2. Return `-1` and `ErrNotFound` when the value is absent.
3. Treat a nil slice as containing nothing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Find([]int{4, 7, 9}, 7)
Output: 1, nil
```

**Example 2:**

```
Input:  Find([]int{4, 7}, 5)
Output: -1, ErrNotFound
```

**Example 3:**

```
Input:  Find(nil, 1)
Output: -1, ErrNotFound
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Linear search** | First match wins. |
| 2 | **Index plus error** | Position and success are separate answers. |
| 3 | **Duplicates** | Only the earliest index is reported. |

## Hint

A slice containing the target twice must still report the lower index.

## Validate

```bash
make verify
```
