# Match Through Wrapping

**Level:** middle
**Topic:** 04-error-handling

## Context

An HTTP handler must answer 404 whenever the failure was ultimately a missing record — however many layers wrapped it on the way up.

## Task

Implement `IsNotFound` in [isnotfound.go](isnotfound.go):

1. Return `true` when `ErrNotFound` appears anywhere in the chain.
2. Return `false` for any other error.
3. Return `false` for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IsNotFound(ErrNotFound)
Output: true
```

**Example 2:**

```
Input:  IsNotFound(fmt.Errorf("load: %w", ErrNotFound))
Output: true
```

**Example 3:**

```
Input:  IsNotFound(ErrDenied)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **errors.Is** | Walks the whole chain, not just the outermost error. |
| 2 | **Equality is not enough** | `err == ErrNotFound` fails on a wrapped error. |
| 3 | **Nil handling** | `errors.Is(nil, target)` is false for a non-nil target. |

## Hint

One of the test cases wraps the sentinel twice — a direct comparison cannot see through that.

## Validate

```bash
make verify
```
