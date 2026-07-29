# Delete Head Ignores New Head

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Removing the head means returning `head.Next` as the new head. Returning the
original `head` keeps the deleted node.

## Task

Fix [delhead.go](delhead.go) so the first node is actually removed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RemoveFirst(1->2->3)
Output: 2->3
```

**Example 2:**

```
Input:  RemoveFirst(1)
Output: nil
```

**Example 3:**

```
Input:  RemoveFirst(1->2)
Output: 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **New head on head-delete** | Return `head.Next`. |
| 2 | **Caller reassigns** | The returned head replaces the old one. |
| 3 | **Single node** | Returns nil. |

## Hint

Return the next node: `return head.Next`.

## Validate

```bash
make verify
```
