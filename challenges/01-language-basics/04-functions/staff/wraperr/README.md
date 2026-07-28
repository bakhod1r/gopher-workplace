# Defer Wraps the Snapshot

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

`defer f(err)` evaluates `err` at the defer statement — before the body assigns
it — so the closure receives nil. To wrap the FINAL error, read the named return
`err` inside the closure body instead of taking it as an argument.

## Task

Fix [wraperr.go](wraperr.go) so the returned error is wrapped.

Do **not** change the function signature or the tests.

## Examples

```go
Do(true)  // => "do: boom"
Do(false) // => nil
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Defer argument snapshot** | Args evaluate at defer-time (err is nil). |
| 2 | **Named-return wrapping** | Read `err` in the body at return-time. |
| 3 | **%w wrapping** | `fmt.Errorf("...: %w", err)` chains the cause. |

## Hint

Drop the parameter and read the named return in the body: `defer func(){ if err != nil { err = fmt.Errorf("do: %w", err) } }()`.

## Validate

```bash
make verify
```
