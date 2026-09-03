# Is Before Unwrap

## Intuition

At each level `errors.Is` compares values, then asks the error's own `Is`, then unwraps. A custom `Is` adds a matching rule; it never removes the chain behind it.

## Approach

1. Format the code and cause in `Error`.
2. Compare codes in `Is` after a type assertion.
3. Return the cause from `Unwrap`.

## Solution

```go
// Error:
return fmt.Sprintf("code %d: %s", e.Code, e.Cause.Error())

// Is:
t, ok := target.(*CodedError)
return ok && t.Code == e.Code

// Unwrap:
return e.Cause
```

## Walkthrough

Matching `ErrBase` fails the custom `Is` at the top level, then succeeds after `errors.Is` unwraps to the cause.

## Pitfalls

- Omitting `Unwrap`, so a custom `Is` hides the cause.
- Returning true from `Is` for any target, swallowing unrelated queries.
- Comparing causes inside `Is`, duplicating what unwrapping already does.
