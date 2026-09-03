# Extract A Typed Error

**Level:** middle
**Topic:** 04-error-handling

## Context

Transport failures carry a status code. The router needs that number, but only when the failure really was an HTTP error.

## Task

Implement `CodeOf` in [httpcode.go](httpcode.go):

1. Return the `Code` of a `*HTTPError` found anywhere in the chain, and `true`.
2. Return `0, false` for any other error.
3. Return `0, false` for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CodeOf(&HTTPError{Code: 404})
Output: 404, true
```

**Example 2:**

```
Input:  CodeOf(fmt.Errorf("get: %w", &HTTPError{Code: 500}))
Output: 500, true
```

**Example 3:**

```
Input:  CodeOf(errors.New("boom"))
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **errors.As** | Finds the first error in the chain of a given type. |
| 2 | **Pointer receivers** | `*HTTPError` implements `error`; the target is `**HTTPError`. |
| 3 | **Type assertion vs As** | A plain assertion cannot see through wrapping. |

## Hint

`errors.As` takes a pointer to the variable you want filled in — declare the target first.

## Validate

```bash
make verify
```
