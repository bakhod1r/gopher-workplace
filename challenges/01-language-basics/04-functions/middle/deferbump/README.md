# Defer Adjusts Result

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

A deferred closure can read and rewrite a named return value after the
`return` statement has set it — the classic post-processing hook.

## Task

Implement `Compute` in [deferbump.go](deferbump.go): set `n = 10`, defer a closure that does `n *= 2`, then `return`.

Do **not** change the function signature or the tests.

## Examples

```go
Compute() // => 20
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Defer runs post-return** | The named result is already 10 when the defer fires. |
| 2 | **Mutating named result** | `n *= 2` changes what the caller receives. |
| 3 | **Order** | Set, schedule, return, then defer mutates. |

## Hint

Named return `n`; `n = 10`; `defer func(){ n *= 2 }()`; `return`.

## Validate

```bash
make verify
```
