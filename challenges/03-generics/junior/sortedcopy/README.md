# Sorted Copy

**Level:** junior  
**Topic:** 03-generics

## Context

A report needs its rows in order, but the caller's slice is shared with other goroutines and must stay as it is.

## Task

Implement the stub(s) in [sortedcopy.go](sortedcopy.go):

1. Implement `Sorted`, returning a **new** sorted slice.
2. Leave the input slice unmodified.
3. Return an empty (non-nil) slice for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Sorted([]int{3, 1, 2})
Output: []int{1, 2, 3}
```

**Example 2:**

```
Input:  Sorted([]string{"b", "a"})
Output: []string{"a", "b"}
```

**Example 3:**

```
Input:  Sorted([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Sorting with `<`** | An insertion sort needs only `<`, which is exactly what `cmp.Ordered` promises. |
| 3 | **Copy before mutating** | Reused from language basics: sorting in place would reorder the caller's slice. |

## Hint

`make` + `copy`, then insertion-sort the copy with `<`.

## Validate

```bash
make verify
```
