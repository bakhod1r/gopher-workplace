# Behavioural Interface

## Intuition

Coupling to a concrete error type ties the caller to one package. Coupling to a method — "anything that can tell me whether it is temporary" — lets unrelated packages participate.

## Approach

1. Declare `var t interface{ Temporary() bool }`.
2. Call `errors.As(err, &t)`.
3. Return `t.Temporary()` on a match, false otherwise.

## Solution

```go
var t interface{ Temporary() bool }
if errors.As(err, &t) {
	return t.Temporary()
}
return false
```

## Walkthrough

For a wrapped permanent error `errors.As` finds the `*NetError`, but `Temporary()` returns false, so the retry does not happen.

## Pitfalls

- Returning true merely because the method exists.
- Type-asserting to `*NetError`, which excludes other packages' errors and misses wrapping.
- Forgetting the `errors` import when the stub does not use it.
