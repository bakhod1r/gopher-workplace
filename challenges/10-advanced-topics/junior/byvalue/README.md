# Small Values Stay On The Stack

**Level:** junior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A geometry package returns `*Point` from every helper "to avoid copying". Each helper costs an allocation and the copies it avoided were sixteen bytes.

## Task

Implement [byvalue.go](byvalue.go):

1. Return `p` with both coordinates multiplied by `f`.
2. The caller's `Point` must not change.
3. Zero allocations.

Replace the stub body in [byvalue.go](byvalue.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Scale(Point{2,3}, 2)
Output: {4 6}
```

**Example 2:**

```
Input:  p := Point{2,3}; Scale(p, 5)
Output: p is still {2 3}
```

_Explanation:_ Structs are passed by value.

**Example 3:**

```
Input:  Scale(Point{1,1}, 0)
Output: {0 0}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Struct value semantics** | Passing and returning a struct copies its fields. |
| 2 | **Stack allocation** | A value that never escapes needs no heap. |
| 3 | **Copy cost vs allocation cost** | Copying two ints is cheaper than allocating one pointer. |

## Hint

The parameter is already your own copy. Mutate it and hand it back.

## Validate

```bash
make verify
```
