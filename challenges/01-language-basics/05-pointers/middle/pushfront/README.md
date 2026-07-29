# Prepend to List

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Prepending to a singly linked list creates a new node whose Next is the old
head, and returns it as the new head.

## Task

Implement `PushFront` in [pushfront.go](pushfront.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PushFront(1->2, 0)
Output: 0->1->2
```

_Explanation:_ A new node points at the old head, becoming the new head.

**Example 2:**

```
Input:  PushFront(nil, 5)
Output: 5
```

_Explanation:_ Prepending to an empty list yields a single node.

**Example 3:**

```
Input:  PushFront(9, 7)
Output: 7->9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Node construction** | `&Node{Val: v, Next: head}`. |
| 2 | **New head returned** | The caller must reassign. |
| 3 | **nil-safe** | Prepending to nil gives a single node. |

## Hint

`return &Node{Val: v, Next: head}`.

## Validate

```bash
make verify
```
