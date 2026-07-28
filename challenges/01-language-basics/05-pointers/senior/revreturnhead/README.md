# Reverse Returns Old Head

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

After reversal, the original `head` is the LAST node (its Next is nil). The new
head is `prev`. Returning `head` yields a one-element list.

## Task

Fix the return in [revreturnhead.go](revreturnhead.go).

Do **not** change the function signature or the tests.

## Examples

```go
Reverse(1->2->3) // => 3->2->1 (return prev)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **New head is prev** | The loop leaves prev at the front. |
| 2 | **Old head becomes tail** | Its Next is now nil. |
| 3 | **Return the right node** | `return prev`. |

## Hint

Return the new head: `return prev`.

## Validate

```bash
make verify
```
