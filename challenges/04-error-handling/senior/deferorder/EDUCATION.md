# Deferred Cleanup Order

## Intuition

Release order matters because later resources often depend on earlier ones. Reversing the list reproduces `defer`'s discipline when the closers arrive as data rather than as statements.

## Approach

1. Loop from the last index down to zero.
2. Collect non-nil results.
3. Return `errors.Join` of them.

## Solution

```go
var errs []error
for i := len(closers) - 1; i >= 0; i-- {
	if err := closers[i](); err != nil {
		errs = append(errs, err)
	}
}
return errors.Join(errs...)
```

## Walkthrough

Three closers run c, b, a; the failing first and third still both execute, so the join reports two failures.

## Pitfalls

- Iterating forwards, releasing in acquisition order.
- Returning at the first failure, stranding remaining resources.
- Using `defer` inside the loop, which delays every close to function exit.
