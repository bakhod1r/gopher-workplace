# Deep Copy a List

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A deep copy allocates a NEW node for each original node, so the two lists share
no memory.

## Task

Implement `Copy` in [copylist.go](copylist.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c := Copy(1->2->3)
Output: independent 1->2->3
```

_Explanation:_ Nodes are freshly allocated, not shared.

**Example 2:**

```
Input:  Copy(nil)
Output: nil
```

**Example 3:**

```
Input:  c := Copy(orig); c.Val = 9
Output: orig.Val unchanged
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **New node per node** | Allocate fresh `&Node{}`. |
| 2 | **Copy the value** | `Val` is copied; Next is a new copy. |
| 3 | **nil terminator** | Recurse/loop until nil. |

## Hint

`if head == nil { return nil }; return &Node{Val: head.Val, Next: Copy(head.Next)}`.

## Validate

```bash
make verify
```
