# Detached Audit Write

## Intuition

A context bundles two unrelated things: request-scoped values and a cancellation signal. Audit needs the first and must ignore the second. `WithoutCancel` splits them cleanly, where copying values by hand would be error-prone.

## Approach

1. Comma-ok read the actor from `ctx.Value(auditKey{})`.
2. Derive `auditCtx` with `context.WithoutCancel(ctx)`.
3. Call `write(auditCtx, actor)` and return the actor with its error.

## Solution

```go
func Record(ctx context.Context, write func(ctx context.Context, actor string) error) (string, error) {
	actor, _ := ctx.Value(auditKey{}).(string)

	auditCtx := context.WithoutCancel(ctx)
	if err := write(auditCtx, actor); err != nil {
		return actor, err
	}
	return actor, nil
}
```

## Walkthrough

The client hangs up, so the request context is `Canceled`. `WithoutCancel` returns a context whose `Err()` is nil but whose `Value(auditKey{})` still walks up into the request's frames, so `write` sees actor `u2` and a live context, and the record lands.

## Pitfalls

- Passing the original `ctx` to `write`, so the audit write is skipped exactly when it matters most.
- Using `context.Background()` instead — that drops the trace values too.
- Detaching long-running work: without a deadline of its own it can outlive the process's shutdown.
