# Merge Drops the Remainder

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

When one list is exhausted, the other may still have nodes. The merge must link
the remaining list to the tail, or those nodes are lost.

## Task

Fix [mergelost.go](mergelost.go) by attaching the leftover list.

Do **not** change the function signature or the tests.

## Examples

```go
Merge(1->2, 3->4->5) // => 1->2->3->4->5
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Leftover handling** | One list may still have nodes. |
| 2 | **Attach remainder** | `tail.Next = a` or `b`. |
| 3 | **Loop exit** | The loop ends when either is nil. |

## Hint

After the loop, attach whichever remains: `if a != nil { tail.Next = a } else { tail.Next = b }`.

## Validate

```bash
make verify
```
