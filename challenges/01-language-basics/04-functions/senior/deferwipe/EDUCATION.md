# Deferred mutation of named returns

## Intuition

Because defers run after the return value is assigned, a deferred closure that writes the named result overrides whatever the body computed.

## Approach

1. A deferred `result = 0` runs after the body and clobbers the named return.
2. Remove the erroneous defer.

## Solution

```go
func Compute(a, b int) (result int) {
	result = a * b
	return
}
```

## Walkthrough

The body computes 42, then the deferred assignment resets `result` to 0 just before return. Deleting the defer preserves the computed value.

## Pitfalls

- Deferred writes to a named return WIN over the body's assignment.
- Only mutate the result in a defer when you intend to (e.g. wrapping errors).
