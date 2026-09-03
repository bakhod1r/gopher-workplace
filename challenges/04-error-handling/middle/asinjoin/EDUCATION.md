# Find A Type Inside A Join

## Intuition

An error value is a tree once joins are involved: chains go down, joins go sideways. `errors.As` walks the whole tree, so hand-written traversals are both longer and easier to get wrong.

## Approach

1. Declare `var fe *FieldError`.
2. Call `errors.As(err, &fe)`.
3. Return `fe, true` on success, `nil, false` otherwise.

## Solution

```go
var fe *FieldError
if errors.As(err, &fe) {
	return fe, true
}
return nil, false
```

## Walkthrough

For `errors.Join(email, age)` the search visits branches in order and stops at `email`, the first match.

## Pitfalls

- Unwrapping manually and missing the joined branches.
- Type-asserting the joined error itself, which is not a `*FieldError`.
- Returning a non-nil pointer alongside `false`.
