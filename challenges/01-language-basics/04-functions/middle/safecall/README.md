# Recover From Panic

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _defer_

## Context

`recover` inside a deferred function stops a panic from propagating. It only
works when called directly from a deferred call.

## Task

Implement `SafeInvoke` in [safecall.go](safecall.go): defer a recover that sets the named result, then call `f`.

Do **not** change the function signature or the tests.

## Examples

```go
SafeInvoke(func(){})            // => false
SafeInvoke(func(){ panic(1) }) // => true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **defer + recover** | recover must run inside a deferred function. |
| 2 | **Named result set in defer** | The deferred closure flips `recovered`. |
| 3 | **Panic unwinding** | recover halts the unwind at this frame. |

## Hint

`defer func(){ if r := recover(); r != nil { recovered = true } }()` then `f()`.

## Validate

```bash
make verify
```
