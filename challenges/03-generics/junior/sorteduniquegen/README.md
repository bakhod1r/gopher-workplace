# Sorted Unique

**Level:** junior  
**Topic:** 03-generics

## Context

A tag cloud shows each tag once, in a stable alphabetical order, from a stream that repeats tags freely.

## Task

Implement the stub(s) in [sorteduniquegen.go](sorteduniquegen.go):

1. Implement `SortedUnique`, returning the distinct elements in ascending order.
2. Return an empty (non-nil) slice for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SortedUnique([]int{3, 1, 3})
Output: []int{1, 3}
```

**Example 2:**

```
Input:  SortedUnique([]string{"b", "a", "b"})
Output: []string{"a", "b"}
```

**Example 3:**

```
Input:  SortedUnique([]int{})
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Ordered implies comparable** | `T` can be a map key here precisely because ordered types are comparable. |
| 2 | **Combining two passes** | Deduplicate with a set, then sort — each step does one thing. |
| 3 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |

## Hint

Reuse the `seen` map idea, then sort the result.

## Validate

```bash
make verify
```
