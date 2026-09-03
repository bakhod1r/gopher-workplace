# Reserve Once, Fill Later

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Sometimes you learn the final size after you have already started filling a slice — a header parsed, a count read off the wire. Reserving the room in one deliberate step beats letting `append` discover it the expensive way.

## Task

Implement `GrowTo` in [growonce.go](growonce.go):

1. Return a slice with the same length and contents as `s` and capacity at least `n`.
2. Return `s` untouched, with no allocation, when its capacity already suffices.
3. Otherwise allocate exactly once, and do not alias `s`.

## Examples

**Example 1:**

```
Input:  GrowTo([1 2 3], 100)
Output: [1 2 3] with cap >= 100
```

**Example 2:**

```
Input:  GrowTo(make([]int, 3, 50), 10)
Output: the same slice, cap 50, no allocation
```

**Example 3:**

```
Input:  GrowTo(nil, 8)
Output: [] with cap >= 8
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Capacity is reserved, not used** | `len` stays put; the extra room only removes future regrows. |
| 2 | **`make(T, len(s), n)` then `copy`** | Preserves the length while widening the capacity. |
| 3 | **Growing detaches** | After a regrow the new array is independent, which callers may or may not expect. |

## Topics used again

`cap`, `copy`, `make` with length and capacity.

## Hint

One comparison decides everything; the growth path is three lines.

## Validate

```bash
make verify
```
