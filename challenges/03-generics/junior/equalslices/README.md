# Slice Equality

**Level:** junior  
**Topic:** 03-generics

## Context

A cache layer must know whether a freshly fetched slice differs from the stored one before writing an update.

## Task

Implement the stub(s) in [equalslices.go](equalslices.go):

1. Implement `Equal`, reporting whether `a` and `b` hold the same elements in the same order.
2. Slices of different lengths are never equal; two empty slices are equal.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Equal([]int{1, 2}, []int{1, 2})
Output: true
```

**Example 2:**

```
Input:  Equal([]int{1, 2}, []int{2, 1})
Output: false
```

**Example 3:**

```
Input:  Equal([]int{}, []int{})
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |
| 2 | **Slices are not comparable** | You cannot write `a == b` for slices — even with `comparable`, that constrains `T`, not `[]T`. |
| 3 | **Length guard first** | Reused from language basics: checking `len` first avoids an out-of-range index. |

## Hint

`comparable` applies to `T`, so you compare `a[i]` and `b[i]`, never `a` and `b`.

## Validate

```bash
make verify
```
