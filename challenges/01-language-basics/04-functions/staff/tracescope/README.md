# Defer Scope in a Loop

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

A `defer` runs when its ENCLOSING FUNCTION returns, not at the end of a loop
iteration. Deferring inside the loop body schedules every "end" for function
exit, producing [start0 start1 end1 end0]. Scope the defer to the iteration by
wrapping the body in a function literal.

## Task

Fix [tracescope.go](tracescope.go) so each item's end is logged before the next start.

Do **not** change the function signature or the tests.

## Examples

```go
Trace(2) // => [start0 end0 start1 end1]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Defer is function-scoped** | It fires at the enclosing function's return. |
| 2 | **Per-iteration wrapper** | An inner func literal makes defer fire per iteration. |
| 3 | **LIFO at exit** | Loop-level defers unwind in reverse at the end. |

## Hint

Wrap the body in an immediately-invoked function so its defer runs per iteration: `func(k int){ log = append(log, ...start); defer func(){ log = append(log, ...end) }() }(i)`.

## Validate

```bash
make verify
```
