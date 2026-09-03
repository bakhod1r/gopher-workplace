# Root Cause

**Level:** middle
**Topic:** 04-error-handling

## Context

A crash report shows the deepest error in the chain, the one that actually happened, with all the annotation layers stripped away.

## Task

Implement `Root` in [rootcause.go](rootcause.go):

1. Unwrap repeatedly and return the last error in the chain.
2. Return `err` itself when it wraps nothing.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Root(fmt.Errorf("a: %w", ErrBase))
Output: ErrBase
```

**Example 2:**

```
Input:  Root(ErrBase)
Output: ErrBase
```

**Example 3:**

```
Input:  Root(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Iterative unwrapping** | Loop until `Unwrap` returns nil. |
| 2 | **Loop invariant** | Keep the last non-nil error seen. |
| 3 | **Chain depth** | Any number of layers reduces to one root. |

## Hint

Unwrap in a loop, but return the error you had *before* the unwrap returned nil.

## Validate

```bash
make verify
```
