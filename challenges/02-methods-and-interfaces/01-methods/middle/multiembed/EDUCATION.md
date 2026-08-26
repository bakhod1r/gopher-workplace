# Resolving Embedded Method Collisions

## Intuition

Go does not do implicit resolution (like C++'s diamond problem or Python's MRO).
If two embedded structs have the same method, the outer struct *has no* promoted
method of that name. You must define it yourself and explicitly call the one
you want.

## Approach

1. Route the call to the `B` field.

## Solution

```go
func (c Collision) Name() string {
	return c.B.Name()
}
```

## Walkthrough

- `c.Name()` is called.
- It accesses the embedded `B` struct: `c.B`.
- It calls `Name()` on it, returning `"B"`.

## Pitfalls

- Trying to call `c.Name()` (recursive).
- Returning `"B"` as a string literal — misses the point of routing to the
  embedded type.
