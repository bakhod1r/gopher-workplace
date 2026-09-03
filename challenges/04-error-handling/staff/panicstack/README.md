# Stack At The Panic

**Level:** staff
**Topic:** 04-error-handling

## Context

A recovered panic without a stack is nearly useless in a postmortem. The boundary captures the stack while it is still live.

## Task

Implement `Trace` in [panicstack.go](panicstack.go):

1. Return nil when `f` does not panic.
2. Return an error whose message contains the panic value and a stack snippet.
3. Capture the stack inside the deferred function, not after returning.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Trace(func() {})
Output: nil
```

**Example 2:**

```
Input:  Trace(func() { panic("boom") })
Output: contains "boom"
```

**Example 3:**

```
Input:  …
Output: contains a function name
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **runtime.Stack** | Fills a buffer with the current goroutine's stack. |
| 2 | **Capture timing** | The stack is gone once the deferred call returns. |
| 3 | **Bounded buffers** | A fixed buffer keeps the cost predictable. |

## Hint

`runtime.Stack(buf, false)` returns how many bytes it wrote — slice the buffer to that length.

## Validate

```bash
make verify
```
