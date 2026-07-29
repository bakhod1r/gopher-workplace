# Delete First Match

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Deleting from a linked list means re-pointing the predecessor's Next past the
target. Deleting the head returns a new head.

## Task

Implement `Delete` in [dellist.go](dellist.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Delete(1->2->3, 2)
Output: 1->3
```

**Example 2:**

```
Input:  Delete(1->2->3, 1)
Output: 2->3
```

_Explanation:_ Deleting the head returns `head.Next`.

**Example 3:**

```
Input:  Delete(1->2->3, 9)
Output: 1->2->3
```

_Explanation:_ Missing value leaves the list unchanged.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Head special case** | Deleting head returns `head.Next`. |
| 2 | **Predecessor relink** | `prev.Next = cur.Next`. |
| 3 | **First match only** | Stop after removing one. |

## Hint

If `head.Val == v` return `head.Next`; else walk with `prev`, and when `prev.Next.Val == v` set `prev.Next = prev.Next.Next`.

## Validate

```bash
make verify
```
