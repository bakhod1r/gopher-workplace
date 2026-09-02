# Zip With

**Level:** junior  
**Topic:** 03-generics

## Context

Two parallel slices hold names and scores. A report combines them position by position, and the slices are not always the same length.

## Task

Implement the stub(s) in [zipwithgen.go](zipwithgen.go):

1. Implement `ZipWith`, combining `a[i]` and `b[i]` with `f` for each valid position.
2. Stop at the shorter of the two slices.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ZipWith([]int{1, 2}, []int{10, 20}, add)
Output: []int{11, 22}
```

**Example 2:**

```
Input:  ZipWith([]int{1, 2, 3}, []int{10}, add)
Output: []int{11}
```

**Example 3:**

```
Input:  ZipWith([]int{1}, []int{}, add)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Three type parameters** | `T`, `U`, and `R` let the two inputs and the result all differ. |
| 2 | **Index-based loops** | `range` over one slice would not bound the other — an explicit counter is safer here. |
| 3 | **Functions as values** | Reused from language basics: a `func(T) U` parameter is an ordinary value. |

## Hint

Compute the shorter length first, then loop by index.

## Validate

```bash
make verify
```
