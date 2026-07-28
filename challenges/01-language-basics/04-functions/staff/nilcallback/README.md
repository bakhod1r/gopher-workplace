# Optional Callback Nil Check

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A nil function value is callable syntactically but panics when invoked. Optional
callbacks must be nil-checked before the call.

## Task

Fix [nilcallback.go](nilcallback.go) so a nil hook is skipped (element added unchanged).

Do **not** change the function signature or the tests.

## Examples

```go
Process([1 2 3], nil)          // => 6
Process([1 2 3], x*10)         // => 60
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Nil function values** | Calling a nil func panics. |
| 2 | **Optional hooks** | Guard with `if hook != nil`. |
| 3 | **Default behaviour** | Fall back to the identity when nil. |

## Hint

Guard the call: `if hook != nil { total += hook(v) } else { total += v }`.

## Validate

```bash
make verify
```
