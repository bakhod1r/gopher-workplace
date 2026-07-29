# Dedup Sorted List

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

In a sorted list, duplicates are adjacent. Skip a node by pointing the current
node's Next past any equal successor.

## Task

Implement `Dedup` in [deduplist.go](deduplist.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Dedup(1->1->2->3->3)
Output: 1->2->3
```

_Explanation:_ Adjacent duplicates in a sorted list are removed.

**Example 2:**

```
Input:  Dedup(nil)
Output: nil
```

**Example 3:**

```
Input:  Dedup(1->2->3)
Output: 1->2->3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Adjacent duplicates** | Equal values are neighbours. |
| 2 | **Skip by relinking** | `cur.Next = cur.Next.Next`. |
| 3 | **Advance carefully** | Only move on when not skipping. |

## Hint

Walk `cur`; while `cur != nil && cur.Next != nil`: if equal, `cur.Next = cur.Next.Next`, else `cur = cur.Next`.

## Validate

```bash
make verify
```
