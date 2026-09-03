# Cancellation Cause Reporting

## Intuition

`ctx.Err()` answers *whether* a context is finished; `context.Cause` answers *why*. Cancelling with a cause lets one cancellation channel carry a whole taxonomy of reasons to whoever inspects it later.

## Approach

1. Derive `runCtx` with `WithCancelCause` and `defer cancel(nil)`.
2. Before writing each row, return `context.Cause(runCtx)` if the context is finished.
3. At the quota, `cancel(ErrQuotaExceeded)` and return the cause.
4. Return `nil` when every row was written within the quota.

## Solution

```go
func Export(ctx context.Context, rows []string, quota int) (int, error) {
	runCtx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	written := 0
	for _, row := range rows {
		if err := runCtx.Err(); err != nil {
			return written, context.Cause(runCtx)
		}
		if written >= quota {
			cancel(ErrQuotaExceeded)
			return written, context.Cause(runCtx)
		}
		_ = row
		written++
	}
	if written >= quota && len(rows) > quota {
		cancel(ErrQuotaExceeded)
		return written, context.Cause(runCtx)
	}
	return written, nil
}
```

## Walkthrough

Five rows and a quota of three: rows 1-3 are written; on the fourth iteration the quota check fires, `cancel(ErrQuotaExceeded)` records the cause, and `context.Cause` hands it back. If instead the client had hung up, the first `runCtx.Err()` check would fire and `Cause` would report `context.Canceled` inherited from the parent.

## Pitfalls

- Returning `runCtx.Err()` on the quota path — that is `context.Canceled` and the reason is lost.
- Forgetting `defer cancel(nil)`, which leaks the cancel registration on the success path.
- Overwriting an inherited cause: the first cancellation wins, later `cancel` calls are ignored.
