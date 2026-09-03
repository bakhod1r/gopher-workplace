# Cleanup That Panics

## Intuition

Cleanup runs on the failure path, which is exactly when it is least tested. Isolating it in its own recovered scope keeps a broken cleanup from destroying the report of the original failure.

## Approach

1. Defer a recovering closure that also runs the cleanup in its own recovered scope.
2. Convert any recovered value into a wrapper around `ErrPanic`.
3. Return `work()` normally.

## Solution

```go
defer func() {
	if r := recover(); r != nil {
		err = fmt.Errorf("%v: %w", r, ErrPanic)
	}
	func() {
		defer func() {
			if r := recover(); r != nil && err == nil {
				err = fmt.Errorf("%v: %w", r, ErrPanic)
			}
		}()
		cleanup()
	}()
}()
return work()
```

## Walkthrough

When the work panics, the outer recovery sets the error and the inner scope still runs the cleanup; a panic there cannot overwrite the already-recorded failure.

## Pitfalls

- Calling `cleanup` without its own recovery, so its panic escapes.
- Letting the cleanup panic overwrite the work failure.
- Running cleanup before the recovery, so a work panic skips it.
