# Opaque Error Surface

## Intuition

Exported sentinels and types become API you can never change. Exporting predicates instead keeps the representation free while giving callers exactly the decisions they need.

## Approach

1. Implement each predicate as an `errors.Is` against its internal cause.
2. Keep the causes unexported in real code.

## Solution

```go
// IsRetryable:
return errors.Is(err, ErrTransient)

// IsRejected:
return errors.Is(err, ErrInvalid)
```

## Walkthrough

A wrapped `ErrTransient` answers true to retryable and false to rejected — the two questions are independent.

## Pitfalls

- Exporting the sentinels as the primary API, freezing the representation.
- Making one predicate the negation of the other; unknown errors are neither.
- Using `==` and breaking on wrapped errors.
