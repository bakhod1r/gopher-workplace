# Peel One Layer

**Level:** middle
**Topic:** 04-error-handling

## Context

A debug endpoint shows the immediate cause of a failure, one level down, so the operator can see what the current layer was reacting to.

## Task

Implement `Cause` in [unwrapone.go](unwrapone.go):

1. Return the error directly wrapped by `err`.
2. Return nil when `err` does not wrap anything.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Cause(fmt.Errorf("a: %w", ErrBase))
Output: ErrBase
```

**Example 2:**

```
Input:  Cause(ErrBase)
Output: nil
```

**Example 3:**

```
Input:  Cause(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **errors.Unwrap** | Removes exactly one layer. |
| 2 | **Unwrappable errors** | Only errors with an `Unwrap` method have a cause. |
| 3 | **One level vs whole chain** | Unwrap is not recursive. |

## Hint

A doubly-wrapped error unwraps to another wrapper, not to the sentinel underneath.

## Validate

```bash
make verify
```
