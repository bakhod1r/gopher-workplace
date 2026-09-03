# One Type, Both Shapes

## Intuition

Go's unwrap protocol has two shapes and a type may implement only one of them — implementing both is undefined by the standard library. The multi shape covers matching; a named accessor covers the primary.

## Approach

1. Return the primary's message from `Error`.
2. Build a slice with the primary first, then the others.
3. Return the primary from the accessor.

## Solution

```go
// Error:
return m.First.Error()

// Unwrap:
out := make([]error, 0, len(m.Others)+1)
out = append(out, m.First)
return append(out, m.Others...)

// Primary:
return m.First
```

## Walkthrough

`errors.Is` fans out over the returned slice, so both the primary and the secondary match while `Error()` stays a single line.

## Pitfalls

- Implementing both unwrap shapes, whose interaction is unspecified.
- Omitting the primary from the member slice, so it stops matching.
- Returning `m.Others` directly and letting callers mutate the struct's slice.
