# Error Interface

**Level:** junior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A validation layer returns its own error type so callers can read the offending field.

## Task

Implement the stub(s) in [erriface.go](erriface.go):

1. Implement `Error` on `ValidationError` so it renders `"<Field>: <Message>"`.
2. Implement `Validate`, which returns a `*ValidationError` when the name is empty and nil otherwise.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  (&ValidationError{Field: "name", Message: "required"}).Error()
Output: "name: required"
```

**Example 2:**

```
Input:  Validate("")
Output: error "name: required"
```

**Example 3:**

```
Input:  Validate("ann")
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`error` is an interface** | Any type with `Error() string` is an error. |
| 2 | **Custom error types** | Carrying structured fields beats string formatting. |
| 3 | **Nil error return** | Reused: returning `nil` for the success path. |

## Hint

Declare the result as `error` and return `nil` on success — never a typed nil pointer.

## Validate

```bash
make verify
```
