# Total Any Slice Of Integers

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A metrics helper sums whichever integer slice a plugin returns. The plugin API is `any`, and a type switch over every width is four near-identical branches.

## Task

Implement [reflectsum.go](reflectsum.go):

1. Total `v` when it is a slice or array of a signed integer kind.
2. Report false for any other kind, including unsigned elements and non-sequences.
3. Accumulate in int64 so wide values cannot overflow.

Replace the stub body in [reflectsum.go](reflectsum.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Sum([]int32{-1,1})
Output: 0, true
```

**Example 2:**

```
Input:  Sum([]myInt{2,3})
Output: 5, true
```

_Explanation:_ A named type's kind is still int.

**Example 3:**

```
Input:  Sum([]uint8{1})
Output: 0, false
```

_Explanation:_ Unsigned is out of scope.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value.Int** | Reads any signed integer kind as an int64. |
| 2 | **Type.Elem** | A slice or array type knows its element type. |
| 3 | **Kind covers named types** | `myInt`'s kind is int, so it is handled without a special case. |
| 4 | **Index into a Value** | `rv.Index(i)` works for both slices and arrays. |

## Hint

Two checks: the container's kind, then the element's.

## Validate

```bash
make verify
```
