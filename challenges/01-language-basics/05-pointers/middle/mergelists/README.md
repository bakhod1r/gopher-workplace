# Merge Two Sorted Lists

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Merging uses a dummy head and a tail pointer, always attaching the smaller
current node and advancing that list.

## Task

Implement `Merge` in [mergelists.go](mergelists.go).

Do **not** change the function signature or the tests.

## Examples

```go
Merge(1->3->5, 2->4->6) // => 1->2->3->4->5->6
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Dummy head** | Simplifies the first attachment. |
| 2 | **Tail pointer** | Append the smaller node. |
| 3 | **Attach remainder** | Link the non-empty list at the end. |

## Hint

Use a `dummy` and `tail`; while both non-nil attach the smaller and advance; then `tail.Next = a or b` (whichever remains); return `dummy.Next`.

## Validate

```bash
make verify
```
