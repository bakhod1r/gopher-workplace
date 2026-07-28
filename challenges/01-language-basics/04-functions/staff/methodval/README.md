# Method Value Binds Receiver

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A method value like `c.Get` copies the receiver at the moment it is created (for
a value receiver). Re-binding `f = c.Get` after mutating `c` captures the NEW
receiver state, defeating the early snapshot.

## Task

Fix [methodval.go](methodval.go) so the returned function reports the early value.

Do **not** change the function signature or the tests.

## Examples

```go
BoundEarly(7)() // => 7
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Method values** | `c.Get` copies the receiver at bind time. |
| 2 | **Value receiver copy** | Later mutations to `c` don't affect the bound value. |
| 3 | **Re-binding** | Reassigning after mutation captures the new state. |

## Hint

Remove the re-binding `f = c.Get`; keep the method value captured before `c.n = 999`.

## Validate

```bash
make verify
```
