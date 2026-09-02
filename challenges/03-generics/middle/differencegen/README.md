# Difference

**Level:** middle  
**Topic:** 03-generics

## Context

A sync job needs the records present locally but missing remotely, listed once each in local order.

## Task

Implement the stub(s) in [differencegen.go](differencegen.go):

1. Implement `Difference`, returning the distinct elements of `a` absent from `b`.
2. Preserve `a`'s first-seen order and emit each value once.
3. Do not modify either input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Difference([]int{1,2,2,3}, []int{2})
Output: []int{1,3}
```

**Example 2:**

```
Input:  Difference([]int{1}, []int{1})
Output: []int{}
```

**Example 3:**

```
Input:  Difference([]int{}, []int{1})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two sets, two jobs** | One map excludes; the other deduplicates the output. |
| 2 | **Linear beats quadratic** | Probing a set is O(1); scanning `b` per element would be O(n·m). |
| 3 | **Stable order** | Appending in traversal order preserves the input's relative order. |

## Hint

Build the exclusion set from `b` first, then walk `a` once.

## Validate

```bash
make verify
```
