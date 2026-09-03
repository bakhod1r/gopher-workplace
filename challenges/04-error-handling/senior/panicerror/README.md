# Typed Panic Payload

**Level:** senior
**Topic:** 04-error-handling

## Context

An interpreter panics with an error value for control flow and with strings for internal bugs. The boundary must tell them apart.

## Task

Implement `Capture` in [panicerror.go](panicerror.go):

1. Return the recovered value as-is when it is already an `error`.
2. Wrap any other payload as `"panic: <value>"`.
3. Return nil when `f` does not panic.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Capture(func() { panic(ErrStop) })
Output: ErrStop
```

**Example 2:**

```
Input:  Capture(func() { panic("bug") })
Output: "panic: bug"
```

**Example 3:**

```
Input:  Capture(func() {})
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type switch on any** | `recover()` returns `any`. |
| 2 | **Preserving error identity** | An error payload must stay matchable. |
| 3 | **Payload discipline** | Panicking with errors makes recovery precise. |

## Hint

`errors.Is` must still match the sentinel that was panicked — reformatting it would break that.

## Validate

```bash
make verify
```
