# Concat

**Level:** junior  
**Topic:** 03-generics

## Context

Several sources each return part of a result set. The handler merges them without disturbing any source slice.

## Task

Implement the stub(s) in [concatgen.go](concatgen.go):

1. Implement `Concat`, joining the given slices in order into one new slice.
2. Return an empty (non-nil) slice when no slices are given.
3. Do not alias any input slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Concat([]int{1}, []int{2, 3})
Output: []int{1, 2, 3}
```

**Example 2:**

```
Input:  Concat([]string{"a"})
Output: []string{"a"}
```

**Example 3:**

```
Input:  Concat[int]()
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic of slices** | `slices ...[]T` makes each argument a whole slice, collected into a `[][]T`. |
| 2 | **Explicit instantiation** | `Concat[int]()` needs the type argument — with no arguments there is nothing to infer from. |
| 3 | **Fresh backing array** | Reused from language basics: appending onto an input slice could overwrite its spare capacity. |

## Hint

Same shape as `Flatten`, but the groups arrive as variadic arguments.

## Validate

```bash
make verify
```
