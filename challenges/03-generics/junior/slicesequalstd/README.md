# Equal From Stdlib

**Level:** junior  
**Topic:** 03-generics

## Context

A cache invalidator compares the freshly fetched list with the stored one before writing an update.

## Task

Implement the stub(s) in [slicesequalstd.go](slicesequalstd.go):

1. Implement `SameOrder` using `slices.Equal`.
2. A nil slice and an empty slice count as equal.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SameOrder([]int{1, 2}, []int{1, 2})
Output: true
```

**Example 2:**

```
Input:  SameOrder([]int{1}, []int{1, 2})
Output: false
```

**Example 3:**

```
Input:  SameOrder(nil, []int{})
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Equal`** | Compares length then elements — the loop you would have written. |
| 2 | **Slices are not `==`** | Reused from earlier: `a == b` does not compile for slices. |
| 3 | **The `slices` package** | The stdlib ships generic slice helpers — prefer them over hand-rolled loops. |

## Hint

`slices.Equal` already treats nil and empty as equal.

## Validate

```bash
make verify
```
