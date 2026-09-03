# Typed Nil Trap

**Level:** senior
**Topic:** 04-error-handling

## Context

A helper returns a concrete error type from an internal function and assigns it to an `error` variable. The result is non-nil even when nothing failed.

## Task

Implement `Wrap` in [nilinterface.go](nilinterface.go):

1. Return nil when `e` is a nil `*OpError`.
2. Return `e` as an `error` otherwise.
3. Make `Wrap(nil) == nil` true when compared against a nil `error`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Wrap(nil) == nil
Output: true
```

**Example 2:**

```
Input:  Wrap(&OpError{Op: "read"}) == nil
Output: false
```

**Example 3:**

```
Input:  Wrap(&OpError{Op: "read"}).Error()
Output: "read failed"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Typed nil in an interface** | A nil pointer in an interface is not a nil interface. |
| 2 | **Explicit nil check** | Convert to the interface only when the pointer is set. |
| 3 | **Silent failure paths** | Callers' `err != nil` fires on a typed nil. |

## Hint

Returning the pointer directly stores a non-nil type descriptor even when the pointer itself is nil.

## Validate

```bash
make verify
```
