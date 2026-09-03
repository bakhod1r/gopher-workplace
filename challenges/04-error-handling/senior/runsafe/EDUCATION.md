# Runtime Panic Boundary

## Intuition

A recovered panic and a returned error mean different things — one is a bug, the other an expected outcome. Tagging recovered panics with their own sentinel keeps that distinction visible upstream.

## Approach

1. Defer a closure that recovers.
2. Format the payload with `%v` and wrap `ErrRuntime` with `%w`.
3. Return `f()` normally.

## Solution

```go
defer func() {
	if r := recover(); r != nil {
		err = fmt.Errorf("%v: %w", r, ErrRuntime)
	}
}()
return f()
```

## Walkthrough

The index panic's payload renders as `"runtime error: index out of range [3] with length 0"`, and the wrapper keeps `ErrRuntime` matchable alongside it.

## Pitfalls

- Reporting returned errors as panics by wrapping everything.
- Dropping the payload message, leaving an unactionable error.
- Recovering and returning nil, hiding the bug entirely.
