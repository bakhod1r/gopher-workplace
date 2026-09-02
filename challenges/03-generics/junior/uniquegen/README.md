# Unique

**Level:** junior  
**Topic:** 03-generics

## Context

An import job receives the same record more than once. Downstream code needs each value exactly once, in first-seen order.

## Task

Implement the stub(s) in [uniquegen.go](uniquegen.go):

1. Implement `Unique`, dropping later duplicates and keeping the first occurrence of each value.
2. Preserve the order of first occurrences.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Unique([]int{1, 2, 1})
Output: []int{1, 2}
```

**Example 2:**

```
Input:  Unique([]string{"a", "a", "b"})
Output: []string{"a", "b"}
```

**Example 3:**

```
Input:  Unique([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |
| 2 | **Maps with type parameters** | A map key type must be comparable, so `K` needs the `comparable` constraint. |
| 3 | **Set idiom** | Reused from language basics: `map[T]bool` records membership; a missing key reads as `false`. |

## Hint

`comparable` is exactly what a map key type requires — that is why `map[T]bool` compiles here.

## Validate

```bash
make verify
```
