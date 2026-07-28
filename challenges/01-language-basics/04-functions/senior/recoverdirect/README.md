# Recover Must Be Deferred

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

`recover` only stops a panic when called from within a function that is running
because of a `defer` during the panic unwind. Called directly at the top of the
body, `recover` returns nil and the panic still crashes the program.

## Task

Fix the recovery in [recoverdirect.go](recoverdirect.go) so a panicking `f` is caught.

Do **not** change the function signature or the tests.

## Examples

```go
Guard(func(){})            // => false
Guard(func(){ panic(1) })  // => true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **recover placement** | It must run inside a deferred function. |
| 2 | **Panic unwinding** | Defers run during unwind; recover intercepts there. |
| 3 | **Named result in defer** | The deferred closure sets `ok`. |

## Hint

Wrap the recover in a deferred closure and call `f()` after: `defer func(){ if recover() != nil { ok = true } }(); f()`.

## Validate

```bash
make verify
```
