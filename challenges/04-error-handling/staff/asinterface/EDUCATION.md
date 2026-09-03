# Find By Behaviour

## Intuition

Matching on behaviour instead of type is what lets errors from packages that know nothing about each other participate in the same policy.

## Approach

1. Declare `var t interface{ RetryAfter() int }`.
2. Call `errors.As(err, &t)`.
3. Return the delay on success, `0, false` otherwise.

## Solution

```go
var t interface{ RetryAfter() int }
if errors.As(err, &t) {
	return t.RetryAfter(), true
}
return 0, false
```

## Walkthrough

`errors.As` walks the join, finds the `*Throttled` branch satisfying the interface, and assigns it to the target.

## Pitfalls

- Passing the interface value rather than its address.
- Asserting to `*Throttled`, excluding other packages' implementations.
- Walking the tree by hand and missing the joined branches.
