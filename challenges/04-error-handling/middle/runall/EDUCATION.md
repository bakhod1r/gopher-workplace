# Run Every Step

## Intuition

Aborting on the first failure is right when later steps depend on earlier ones. When they are independent, stopping early hides work the operator would have had to redo anyway.

## Approach

1. Declare `var errs []error`.
2. Call each function and append non-nil results.
3. Return `errors.Join(errs...)`.

## Solution

```go
var errs []error
for _, f := range fs {
	if err := f(); err != nil {
		errs = append(errs, err)
	}
}
return errors.Join(errs...)
```

## Walkthrough

With three steps the middle one succeeds and is skipped; the join holds two branches and matches both sentinels.

## Pitfalls

- Returning on the first failure, leaving later steps unrun.
- Appending nil results and relying on `Join` to filter — harmless but obscures intent.
- Returning a non-nil join when nothing failed.
