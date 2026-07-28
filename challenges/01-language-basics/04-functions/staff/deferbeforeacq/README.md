# Defer Before Acquisition

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

A cleanup `defer` must be scheduled AFTER the resource is successfully acquired.
Scheduling it first means the early `return` on the failure path still runs the
release, cleaning up something that was never created.

## Task

Fix [deferbeforeacq.go](deferbeforeacq.go) so nothing is released when nothing is acquired.

Do **not** change the function signature or the tests.

## Examples

```go
Use(false) // => []
Use(true)  // => [open close]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Acquire-then-defer** | Schedule cleanup only after acquisition succeeds. |
| 2 | **Early return path** | A pre-scheduled defer fires on failure too. |
| 3 | **Resource discipline** | Release must pair with a real acquire. |

## Hint

Move the defer below the `if !ok { return }` guard, right after logging "open", so it's only scheduled once the resource exists.

## Validate

```bash
make verify
```
