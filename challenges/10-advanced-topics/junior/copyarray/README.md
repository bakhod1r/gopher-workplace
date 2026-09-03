# Arrays Are Values, Slices Are Not

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A geometry helper takes a fixed-size array and mutates it, expecting the caller to see the change. The caller never does, and the confusion costs an afternoon.

## Task

Implement [copyarray.go](copyarray.go):

1. Return an array whose elements are one greater than `a`'s.
2. The caller's array must be unchanged.
3. No allocations — the array lives on the stack.

Replace the stub body in [copyarray.go](copyarray.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Bump([4]int{1,2,3,4})
Output: [2 3 4 5]
```

**Example 2:**

```
Input:  a := [4]int{1,2,3,4}; Bump(a)
Output: a is still [1 2 3 4]
```

_Explanation:_ The parameter was a copy.

**Example 3:**

```
Input:  Bump([4]int{})
Output: [1 1 1 1]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Array value semantics** | Passing or assigning an array copies every element. |
| 2 | **Array vs slice** | A slice copies a header; an array copies the data. |
| 3 | **Stack allocation** | A small fixed-size value needs no heap. |

## Hint

You may mutate the parameter freely — it is already your copy.

## Validate

```bash
make verify
```
