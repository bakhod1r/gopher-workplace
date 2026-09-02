# Cartesian Product

**Level:** middle  
**Topic:** 03-generics

## Context

A test matrix runs every browser against every screen size, and the order must be stable so failures are comparable between runs.

## Task

Implement the stub(s) in [cartesiangen.go](cartesiangen.go):

1. Implement `Product`, returning every combination with `as` varying slowest.
2. Return an empty (non-nil) result when either input is empty.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Product([]int{1,2}, []string{"a"})
Output: [{1 a} {2 a}]
```

**Example 2:**

```
Input:  Product([]int{1}, []string{"a","b"})
Output: [{1 a} {1 b}]
```

**Example 3:**

```
Input:  Product([]int{}, []string{"a"})
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nested loops define the order** | The outer loop varies slowest — that is what "a-major" means. |
| 2 | **Exact capacity** | `len(as)*len(bs)` is the final length, so one allocation suffices. |
| 3 | **Quadratic growth** | The result is the product of the input sizes; say so before someone passes two big slices. |

## Hint

Outer loop over `as`, inner over `bs`.

## Validate

```bash
make verify
```
