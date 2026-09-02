# Reverse

**Level:** junior  
**Topic:** 03-generics

## Context

A history panel shows the newest entry first, while the store keeps entries oldest-first. The stored slice must stay untouched.

## Task

Implement the stub(s) in [reversegen.go](reversegen.go):

1. Implement `Reverse`, returning a **new** slice with the elements in the opposite order.
2. Leave the input slice unmodified.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Reverse([]int{1, 2, 3})
Output: []int{3, 2, 1}
```

**Example 2:**

```
Input:  Reverse([]string{"a", "b"})
Output: []string{"b", "a"}
```

**Example 3:**

```
Input:  Reverse([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`make` with a type parameter** | `make([]T, 0, len(s))` allocates for an unknown element type. |
| 2 | **Aliasing** | Reused from language basics: writing into `s` would be visible to the caller. |
| 3 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |

## Hint

Count down from `len(s)-1` to `0` and append.

## Validate

```bash
make verify
```
