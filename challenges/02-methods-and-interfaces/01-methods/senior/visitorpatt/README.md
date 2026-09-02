# Visitor Pattern

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A tree knows how to walk itself; the caller supplies what to *do* at each node.
That split is the visitor pattern: traversal lives with the structure, behaviour
arrives as a parameter.

## Task

Implement `Accept` on `*Node` in [visitorpatt.go](visitorpatt.go):

1. Return immediately if the receiver is nil.
2. Call `visitor(n.Val)`.
3. Recurse into `n.Left`, then `n.Right`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  single node {Val: 1}; visitor sums
Output: sum == 1
```

**Example 2:**

```
Input:  root 1 with children 2 and 3
Output: sum == 6
```

**Example 3:**

```
Input:  visitor collecting values, same tree
Output: [1 2 3]  (pre-order: node, left, right)
```

_Explanation:_ visiting before recursing gives pre-order.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nil receiver is legal** | A method on a pointer type may be called with a nil receiver — that is what makes the base case a one-liner. |
| 2 | **Recursive traversal** | Left then right, so the order is deterministic. |
| 3 | **Function as visitor** | `func(int)` replaces the classic visitor interface. |

## Hint

`if n == nil { return }` at the top. Because that check exists, you can recurse
unconditionally: `n.Left.Accept(visitor)` is safe even when `Left` is nil.

## Validate

```bash
make verify
```
