# Isolate A Callback

**Level:** senior
**Topic:** 04-error-handling

## Context

A plugin host runs untrusted callbacks. A misbehaving plugin must fail its own call and leave the host running.

## Task

Implement `SafeCall` in [safecall.go](safecall.go):

1. Return `f`'s error unchanged when it returns normally.
2. Return an error wrapping the recovered value as `"panic: <value>"` when `f` panics.
3. Return `ErrNilFunc` when `f` is nil, without panicking.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SafeCall(okFn)
Output: nil
```

**Example 2:**

```
Input:  SafeCall(panicFn)
Output: "panic: boom"
```

**Example 3:**

```
Input:  SafeCall(nil)
Output: ErrNilFunc
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Recovering a foreign panic** | The payload can be any type. |
| 2 | **Formatting an arbitrary value** | `%v` renders any panic payload. |
| 3 | **Nil function guard** | Calling a nil func is itself a panic. |

## Hint

The nil check can be an explicit guard or handled by the same recovery — but the tests demand a specific sentinel for it.

## Validate

```bash
make verify
```
