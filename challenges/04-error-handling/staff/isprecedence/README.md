# Is Before Unwrap

**Level:** staff
**Topic:** 04-error-handling

## Context

A code-carrying error matches by class through its own `Is`, but its wrapped cause must remain matchable too.

## Task

Implement `CodedError` in [isprecedence.go](isprecedence.go):

1. Give `*CodedError` an `Error() string` of `"code <Code>: <Cause>"`.
2. Give it an `Is(target error) bool` matching another `*CodedError` with the same code.
3. Give it an `Unwrap() error` so the cause still matches.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  errors.Is(e, &CodedError{Code: 7})
Output: true
```

**Example 2:**

```
Input:  errors.Is(e, ErrBase)
Output: true
```

**Example 3:**

```
Input:  errors.Is(e, &CodedError{Code: 9})
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Match order** | `Is` is tried at each level, then the chain continues. |
| 2 | **Coexisting mechanisms** | A custom `Is` does not disable unwrapping. |
| 3 | **Class and cause** | Two independent questions about one error. |

## Hint

A custom `Is` returning false does not end the search — `errors.Is` keeps unwrapping.

## Validate

```bash
make verify
```
