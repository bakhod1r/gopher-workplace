# Attach The Call Site

**Level:** senior
**Topic:** 04-error-handling

## Context

A library records where an error was created so on-call engineers do not have to guess which of five call sites produced it.

## Task

Implement `Here` in [withstack.go](withstack.go):

1. Return nil when `err` is nil.
2. Return an error whose message is `"<file>:<line>: <err>"` for the caller's position.
3. Keep the original error matchable.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Here(ErrBase)
Output: "withstack_test.go:21: base"
```

**Example 2:**

```
Input:  Here(nil)
Output: nil
```

**Example 3:**

```
Input:  errors.Is(Here(ErrBase), ErrBase)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **runtime.Caller** | Skip level 1 reports the caller, not this function. |
| 2 | **filepath.Base** | Only the file name belongs in the message. |
| 3 | **Wrapping with position** | Location is context, not a replacement. |

## Hint

`runtime.Caller(0)` names this function's own file — the caller is one frame up.

## Validate

```bash
make verify
```
