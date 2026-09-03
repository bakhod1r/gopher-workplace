# Chain Length

**Level:** middle
**Topic:** 04-error-handling

## Context

A latency investigation counts how many layers each failure passed through before reaching the handler.

## Task

Implement `Depth` in [wrapdepth.go](wrapdepth.go):

1. Return the number of errors in the chain, counting `err` itself.
2. Return `0` for a nil error.
3. Return `1` for an error that wraps nothing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Depth(nil)
Output: 0
```

**Example 2:**

```
Input:  Depth(ErrBase)
Output: 1
```

**Example 3:**

```
Input:  Depth(fmt.Errorf("a: %w", ErrBase))
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Counting a chain** | Each `Unwrap` step is one link. |
| 2 | **Off-by-one** | The outermost error counts too. |
| 3 | **Nil terminator** | The loop ends when `Unwrap` yields nil. |

## Hint

Increment before unwrapping, not after — the outermost error is a link as well.

## Validate

```bash
make verify
```
