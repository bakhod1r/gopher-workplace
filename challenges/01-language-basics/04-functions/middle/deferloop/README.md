# Defer Inside Loop

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

Defers registered in a loop all fire at function exit, in reverse of
registration — so a loop 0..n-1 unwinds n-1..0.

## Task

Implement `CloseOrder` in [deferloop.go](deferloop.go): loop i from 0 to n-1, deferring an append of i to the named result.

Do **not** change the function signature or the tests.

## Examples

```go
CloseOrder(3) // => [2 1 0]
CloseOrder(1) // => [0]
CloseOrder(0) // => []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Loop-registered defers** | Each iteration schedules one. |
| 2 | **LIFO unwinding** | Later i's run first. |
| 3 | **Per-iteration capture (Go 1.22)** | Each i is distinct. |

## Hint

`for i := 0; i < n; i++ { defer func(v int){ out = append(out, v) }(i) }`.

## Validate

```bash
make verify
```
