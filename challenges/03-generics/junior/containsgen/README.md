# Contains

**Level:** junior  
**Topic:** 03-generics

## Context

A permission check asks whether a role appears in a list. The same scan is needed for IDs, names, and codes.

## Task

Implement the stub(s) in [containsgen.go](containsgen.go):

1. Implement `Contains`, reporting whether any element of `s` equals `v`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Contains([]int{1, 2, 3}, 2)
Output: true
```

**Example 2:**

```
Input:  Contains([]string{"a"}, "b")
Output: false
```

**Example 3:**

```
Input:  Contains([]int{}, 1)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |
| 2 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |
| 3 | **`range` over a slice** | Reused from language basics: `for _, e := range s` walks elements. |

## Hint

`any` is not enough here — `==` needs `comparable`.

## Validate

```bash
make verify
```
