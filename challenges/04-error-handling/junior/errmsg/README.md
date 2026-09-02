# Error Message

**Level:** junior
**Topic:** 04-error-handling

## Context

A log formatter writes one line per step. Successful steps print an empty reason instead of crashing on a nil error.

## Task

Implement `Message` in [errmsg.go](errmsg.go):

1. Return `err.Error()` when `err` is non-nil.
2. Return the empty string when `err` is nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Message(nil)
Output: ""
```

**Example 2:**

```
Input:  Message(errors.New("disk full"))
Output: "disk full"
```

**Example 3:**

```
Input:  Message(ErrTimeout)
Output: "timeout"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The error interface** | `error` has one method: `Error() string`. |
| 2 | **Nil guard** | Calling a method on a nil interface value panics. |
| 3 | **Zero value of string** | `""` is the natural "nothing to report" answer. |

## Hint

The method call is safe only after the nil check passes.

## Validate

```bash
make verify
```
