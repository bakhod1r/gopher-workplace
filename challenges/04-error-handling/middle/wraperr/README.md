# Annotate And Wrap

**Level:** middle
**Topic:** 04-error-handling

## Context

A storage layer returns bare errors that say nothing about which operation failed. The service layer adds the operation name without hiding the original.

## Task

Implement `Wrap` in [wraperr.go](wraperr.go):

1. Return `nil` when `err` is nil.
2. Return an error whose message is `"<op>: <err>"`.
3. Keep the original error reachable by `errors.Is` and `errors.Unwrap`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Wrap("read", ErrDisk)
Output: "read: disk offline"
```

**Example 2:**

```
Input:  Wrap("read", nil)
Output: nil
```

**Example 3:**

```
Input:  errors.Is(Wrap("read", ErrDisk), ErrDisk)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **fmt.Errorf with %w** | `%w` records the wrapped error instead of flattening it. |
| 2 | **Wrapping preserves identity** | `errors.Is` still finds the original sentinel. |
| 3 | **Nil in, nil out** | Wrapping nothing must not manufacture a failure. |

## Hint

`%v` and `%w` produce the same message. Only one of them keeps the chain intact.

## Validate

```bash
make verify
```
