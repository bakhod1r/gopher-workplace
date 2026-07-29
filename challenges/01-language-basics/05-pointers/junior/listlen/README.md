# Linked List Length

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Traversing a linked list follows `Next` pointers until nil.

## Task

Implement `Length` in [listlen.go](listlen.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Length(nil)
Output: 0
```

**Example 2:**

```
Input:  Length(1 -> 2 -> 3)
Output: 3
```

**Example 3:**

```
Input:  Length(&Node{Value: 7})
Output: 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer traversal** | Walk `n = n.Next` until nil. |
| 2 | **nil terminator** | The last Next is nil. |
| 3 | **Counter** | Count each node. |

## Hint

Loop `for n := head; n != nil; n = n.Next { count++ }`.

## Validate

```bash
make verify
```
