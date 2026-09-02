# Flatten

**Level:** junior  
**Topic:** 03-generics

## Context

Paginated results arrive one page at a time. The caller wants a single flat list.

## Task

Implement the stub(s) in [flattengen.go](flattengen.go):

1. Implement `Flatten`, concatenating the groups in order into one slice.
2. Return an empty (non-nil) slice for no groups.
3. Skip nothing: empty groups simply contribute no elements.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Flatten([][]int{{1, 2}, {3}})
Output: []int{1, 2, 3}
```

**Example 2:**

```
Input:  Flatten([][]string{{}, {"a"}})
Output: []string{"a"}
```

**Example 3:**

```
Input:  Flatten([][]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic append** | `append(out, g...)` splices a whole group in one call. |
| 2 | **Two-pass sizing** | Summing the lengths first lets you allocate the exact capacity once. |
| 3 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |

## Hint

Sum the lengths first, then append each group with `...`.

## Validate

```bash
make verify
```
