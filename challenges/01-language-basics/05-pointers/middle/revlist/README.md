# Reverse a Linked List

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Reversing a list re-points each node's Next to its predecessor, tracking the
previous node as you walk.

## Task

Implement `Reverse` in [revlist.go](revlist.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Reverse(1->2->3)
Output: 3->2->1
```

**Example 2:**

```
Input:  Reverse(nil)
Output: nil
```

**Example 3:**

```
Input:  Reverse(5)
Output: 5
```

_Explanation:_ A single node reverses to itself.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Three pointers** | prev, cur, next. |
| 2 | **Re-point Next** | `cur.Next = prev`. |
| 3 | **Advance** | Move prev and cur forward. |

## Hint

Iterate with `prev`; each step: `next := cur.Next; cur.Next = prev; prev = cur; cur = next`; return prev.

## Validate

```bash
make verify
```
