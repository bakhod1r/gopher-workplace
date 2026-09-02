# GC Mark Phase

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A tracing garbage collector starts from roots and marks everything reachable.
Anything left unmarked when the trace finishes is garbage. This puzzle is the
mark half — the traversal that colours the live set.

## Task

Implement `Mark` on `*Object` in [gcmark.go](gcmark.go):

1. Set `o.Marked = true`.
2. Recursively mark everything in `o.Refs`.

**Constraint (staff):** the trace must terminate on a cyclic graph and mark a 100,000-node chain completely.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  o1 with no refs; o1.Mark()
Output: o1.Marked == true
```

**Example 2:**

```
Input:  o1 → o2; o1.Mark()
Output: both marked
```

**Example 3:**

```
Input:  o1 → o2 → o1 (a cycle); o1.Mark()
Output: both marked, and the trace terminates
```

_Explanation:_ an already-marked object must not be traced again, or the cycle never ends.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reachability traversal** | Marking is depth-first over the object graph. |
| 2 | **The mark bit doubles as a visited set** | Checking it before recursing is what makes cycles safe. |
| 3 | **Nil-safe recursion** | A nil entry in `Refs` must not panic. |

## Hint

Return early if `o` is nil **or already marked**. Object graphs have cycles; a
plain recursion without that check overflows the stack.

## Validate

```bash
make verify
```
