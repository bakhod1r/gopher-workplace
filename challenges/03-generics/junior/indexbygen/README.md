# Index By

**Level:** junior  
**Topic:** 03-generics

## Context

A join step needs random access to rows by ID, but the rows arrive as a slice.

## Task

Implement the stub(s) in [indexbygen.go](indexbygen.go):

1. Implement `IndexBy`, returning a map from `key(v)` to the element.
2. On duplicate keys the later element wins.
3. Return an empty (non-nil) map for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IndexBy(rows, idOf)
Output: map from id to row
```

**Example 2:**

```
Input:  IndexBy([]int{1, 11}, mod10)
Output: {1: 11}
```

**Example 3:**

```
Input:  IndexBy([]int{}, mod10)
Output: {}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Building an index** | One pass turns O(n) lookups into O(1) ones. |
| 2 | **Last write wins** | Assigning unconditionally means the final duplicate is what remains. |
| 3 | **Key extraction** | Reused from earlier: a `func(T) K` decouples the key from the traversal. |

## Hint

Assign unconditionally — the last duplicate naturally wins.

## Validate

```bash
make verify
```
