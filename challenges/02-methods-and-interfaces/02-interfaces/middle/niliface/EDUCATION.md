# Nil Interface Trap

## Intuition

An interface value is a pair: dynamic type and dynamic value. `err == nil` is true only when *both* halves are nil. Assigning a nil `*OpError` to an `error` fills the type half, so the result is non-nil even though the pointer is.

## Approach

1. Return the concrete pointer only inside the failure branch.
2. Return the literal `nil` on the success path.
3. `IsNil` is simply `err == nil`.
4. `FailedCount` counts elements where `err != nil`.

## Solution

```go
func Run(fail bool) error {
	if fail {
		return &OpError{Op: "op"}
	}
	return nil
}

func IsNil(err error) bool { return err == nil }

func FailedCount(errs []error) int {
	n := 0
	for _, err := range errs {
		if err != nil {
			n++
		}
	}
	return n
}
```

## Walkthrough

`TestTypedNilIsNotNil` builds the trap by hand: `typed` is a nil `*OpError`, and `iface` holds type `*OpError` with value nil — non-nil as an interface. `Run` must not produce that shape.

## Pitfalls

- `var e *OpError; if fail { e = ... }; return e` — the success path returns a non-nil interface.
- Checking `err.(*OpError) != nil` to work around it instead of fixing the return.
- Assuming the compiler warns about this; it does not.
