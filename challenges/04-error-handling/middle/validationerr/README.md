# Custom Error Type

**Level:** middle
**Topic:** 04-error-handling

## Context

A form rejects individual fields. The rejection must carry the field name in a form the caller can read, not just render.

## Task

Implement `NewValidation` in [validationerr.go](validationerr.go):

1. Give `*ValidationError` an `Error() string` returning `"<Field>: <Reason>"`.
2. Implement `NewValidation` so it returns a `*ValidationError` as an `error`.
3. Keep both fields readable by callers that type-assert.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewValidation("email", "is required").Error()
Output: "email: is required"
```

**Example 2:**

```
Input:  NewValidation("age", "must be positive").Error()
Output: "age: must be positive"
```

**Example 3:**

```
Input:  errors.As(err, &ve)
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Implementing error** | Any type with `Error() string` is an error. |
| 2 | **Structured errors** | Fields carry data a string cannot. |
| 3 | **Pointer receiver** | `*ValidationError` is the type that implements the interface. |

## Hint

Both the method and the constructor are missing — the constructor's return type is `error`, but the value it returns is a pointer.

## Validate

```bash
make verify
```
