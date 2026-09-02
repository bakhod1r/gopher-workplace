# Count Occurrences

**Level:** junior  
**Topic:** 03-generics

## Context

A tally screen reports how often a given status code shows up in a batch of results.

## Task

Implement the stub(s) in [countgen.go](countgen.go):

1. Implement `Count`, returning the number of elements of `s` equal to `v`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Count([]int{1, 2, 1}, 1)
Output: 2
```

**Example 2:**

```
Input:  Count([]string{"a", "a", "a"}, "a")
Output: 3
```

**Example 3:**

```
Input:  Count([]int{1, 2}, 9)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The `comparable` constraint** | `comparable` is what lets you use `==` and `!=` on a type parameter. |
| 2 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |
| 3 | **Accumulator loops** | Reused from language basics: initialise a counter before the loop, not inside it. |

## Hint

Unlike `Contains`, this one must visit every element — no early return.

## Validate

```bash
make verify
```
