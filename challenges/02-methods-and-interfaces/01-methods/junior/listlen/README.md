# List Length

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A linked list needs a `Len` method to report how many elements it holds.

## Task

Implement `Len` on `*Node` in [listlen.go](listlen.go):

1. Count the number of nodes from the receiver to the end.
2. If the receiver is `nil`, return `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Node{1, &Node{2, &Node{3, nil}}}.Len()
Output: 3
```

**Example 2:**

```
Input:  Node{42, nil}.Len()
Output: 1
```

**Example 3:**

```
Input:  (*Node)(nil).Len()
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil receiver** | Go allows calling methods on nil pointers — handle it. |
| 2 | **Pointer receiver** | `*Node` — needed to handle the nil case. |
| 3 | **Linked list traversal** | Walk `Next` until nil. |

## Hint

Check `n == nil` first (nil receiver is valid in Go). Then walk with a counter.

## Validate

```bash
make verify
```
