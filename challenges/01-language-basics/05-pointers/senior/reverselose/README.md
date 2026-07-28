# Reverse Loses the Rest

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

You must save `cur.Next` BEFORE overwriting it. The bug sets `cur.Next = prev`
first, so the following `next := cur.Next` reads `prev`, losing the remainder.

## Task

Fix the statement order in [reverselose.go](reverselose.go).

Do **not** change the function signature or the tests.

## Examples

```go
Reverse(1->2->3) // => 3->2->1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Save before overwrite** | Capture next first. |
| 2 | **Order matters** | Reassigning Next destroys the link. |
| 3 | **Three-pointer reversal** | next, then re-point, then advance. |

## Hint

Save first: `next := cur.Next; cur.Next = prev`.

## Validate

```bash
make verify
```
