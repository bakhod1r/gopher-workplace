# An Error That Costs Nothing To Return

## Intuition

An error carrying no per-call information does not need to be built per call. A package-level sentinel is allocated once at init and returned by pointer forever after.

## Approach

1. Check `n < 0` and return `ErrNegative`.
2. Check `n > MaxCount` and return `ErrTooLarge`.
3. Return nil.

## Solution

```go
import "errors"

// MaxCount is the largest count Validate accepts.
const MaxCount = 1000

// The conditions Validate can report.
var (
	ErrNegative = errors.New("count is negative")
	ErrTooLarge = errors.New("count is too large")
)

// Validate reports whether n is a usable count: it must be non-negative
// and no greater than MaxCount.
//
// The failures are fixed conditions, so they must be reported with the
// package's sentinel errors rather than a freshly formatted one.
//
// Examples:
//
// 	Validate(-1) => ErrNegative
func Validate(n int) error {
	if n < 0 {
		return ErrNegative
	}
	if n > MaxCount {
		return ErrTooLarge
	}
	return nil
}
```

## Walkthrough

`Validate(-1)` returns the interface value pointing at the existing `ErrNegative`. Nothing is constructed, so `AllocsPerRun` reports 0. `fmt.Errorf("count is negative")` would allocate a string and an error struct on every rejection.

## Pitfalls

- Wrapping the sentinel with `fmt.Errorf("%w", ...)` when there is nothing to add.
- Comparing with `==` in the caller is fine for a bare sentinel, but `errors.Is` survives later wrapping.
