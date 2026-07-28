# Run Once

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A boolean captured in a closure gates a one-time action — the essence of
`sync.Once` without the concurrency.

## Task

Implement `Once` in [oncefn.go](oncefn.go).

Do **not** change the function signature or the tests.

## Examples

```go
do := Once(f); do() // runs f
do()                // no-op
do()                // no-op
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Captured flag** | `done bool` remembers prior execution. |
| 2 | **Guard the call** | Run `f` only when `!done`, then set `done`. |
| 3 | **Closure state** | The flag persists across calls. |

## Hint

Capture `done := false`; return `func(){ if !done { done = true; f() } }`.

## Validate

```bash
make verify
```
