# Intersect Slices

**Level:** middle  
**Topic:** 03-generics

## Context

An access check lists the roles a user holds that the resource also requires, in the user's own order.

## Task

Implement the stub(s) in [intersectgen.go](intersectgen.go):

1. Implement `Intersect`, returning the distinct elements of `a` that also occur in `b`.
2. Preserve `a`'s first-seen order and emit each value once.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Intersect([]int{1,2,2,3}, []int{2,3})
Output: []int{2,3}
```

**Example 2:**

```
Input:  Intersect([]int{1}, []int{2})
Output: []int{}
```

**Example 3:**

```
Input:  Intersect([]int{}, []int{1})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Membership set** | Index the smaller side when you can; here the spec fixes `b`. |
| 2 | **Order comes from `a`** | The result's order is `a`'s, which makes the operation non-commutative in output order. |
| 3 | **Stable order** | Appending in traversal order preserves the input's relative order. |

## Hint

Same shape as `Difference` with the membership test inverted.

## Validate

```bash
make verify
```
