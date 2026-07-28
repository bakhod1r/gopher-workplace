# Remove Nth Node Off-by-One

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

To remove index i you need its PREDECESSOR (index i-1) so you can relink past
it. Walking i steps lands on the target itself; walk i-1 steps.

## Task

Fix the walk in [removeatlist.go](removeatlist.go).

Do **not** change the function signature or the tests.

## Examples

```go
RemoveAt(10->20->30->40, 2) // => 10->20->40
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Predecessor needed** | Stop at index i-1. |
| 2 | **Relink past target** | `prev.Next = prev.Next.Next`. |
| 3 | **Off-by-one walk** | `k < i-1`. |

## Hint

Walk to the predecessor: `for k := 0; k < i-1; k++ { prev = prev.Next }`.

## Validate

```bash
make verify
```
