# Append at Tail via Double Pointer

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Appending to the tail must handle the empty-list case by updating the caller's
head pointer — hence `**Node`. Otherwise walk to the last node and link.

## Task

Implement `Append` in [appendtail.go](appendtail.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  var h *Node; Append(&h, 1); Append(&h, 2)
Output: 1->2
```

**Example 2:**

```
Input:  h := &Node{Val:1}; Append(&h, 2)
Output: 1->2
```

**Example 3:**

```
Input:  var h *Node; Append(&h, 9)
Output: 9
```

_Explanation:_ First append sets the head through the double pointer.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Double pointer for empty case** | `*head = node` when list is empty. |
| 2 | **Walk to tail** | Follow Next until it is nil. |
| 3 | **Link the new node** | `last.Next = &Node{...}`. |

## Hint

If `*head == nil { *head = &Node{Val: v}; return }`; else walk to the last node and set its `Next`.

## Validate

```bash
make verify
```
