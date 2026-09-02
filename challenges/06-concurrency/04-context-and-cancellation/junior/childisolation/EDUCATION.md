# A Cancelled Query Must Not Kill the Request

## Intuition

Deriving a context is how you carve out a sub-operation you can abandon on its own. Cancellation only ever flows downward, so aborting the recommendation lookup leaves the request — and its other in-flight work — completely untouched. That asymmetry is what makes per-call contexts safe to hand out.

## Approach

1. Create the request context; `defer` its cancel.
2. Derive the query context and cancel it immediately.
3. `<-queryCtx.Done()` to confirm the child really finished.
4. Return `reqCtx.Err()`.

## Solution

```go
// RequestErrAfterQueryCancel builds the per-request context, derives a query
// context from it, cancels only the query (the handler abandoned that lookup
// and will fall back to the cache), and returns the error the *request*
// context reports.
//
// Examples:
//
//	RequestErrAfterQueryCancel()          => nil
//	the request context stays usable      => true
//	the result is never context.Canceled  => true
func RequestErrAfterQueryCancel() error {
	reqCtx, cancelReq := context.WithCancel(context.Background())
	defer cancelReq()

	queryCtx, cancelQuery := context.WithCancel(reqCtx)
	cancelQuery()
	<-queryCtx.Done()

	return reqCtx.Err()
}
```

## Walkthrough

- `cancelQuery()` closes only the query context's `Done()` and sets its error to `context.Canceled`.
- The request context has no idea this happened: it keeps an open `Done()` channel and a nil `Err()`.
- The handler can therefore go on to read the cache and write a full response.

## Pitfalls

- Returning the *child's* `Err()` by mistake yields `context.Canceled` and inverts the lesson.
- The request context still needs its own `defer cancel()` — the child's cancellation does not release it.
- Do not reuse a cancelled child for the fallback work; derive a fresh context instead.
