# Filter

**Level:** junior  
**Topic:** 03-generics

## Context

A dashboard shows only the rows a user is allowed to see. The predicate changes per role; the traversal does not.

## Task

Implement the stub(s) in [filtergen.go](filtergen.go):

1. Implement `Filter`, keeping the elements for which `keep(e)` is true.
2. Preserve the original order and return a new slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Filter([]int{1, 2, 3}, isEven)
Output: []int{2}
```

**Example 2:**

```
Input:  Filter([]string{"", "a"}, nonEmpty)
Output: []string{"a"}
```

**Example 3:**

```
Input:  Filter([]int{}, isEven)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Functions as values** | Reused from language basics: a `func(T) U` parameter is an ordinary value. |
| 2 | **Predicate functions** | A `func(T) bool` parameter is the idiomatic Go way to inject a condition. |
| 3 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |

## Hint

Append only when `keep(e)` is true; never remove from `s` itself.

## Validate

```bash
make verify
```
