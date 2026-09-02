# Cancellation Reaches the Query

## Intuition

A context is a node in a tree. `WithCancel` registers the new context as a child of its parent, so cancelling the parent walks down and cancels every descendant. That is how one `cancel()` in the HTTP server tears down every query, RPC and goroutine the request started — no manual bookkeeping anywhere.

## Approach

1. Create the request context with `WithCancel(context.Background())`.
2. Derive the query context with `WithCancel(reqCtx)` and `defer` its cancel.
3. Call the request's cancel.
4. `<-queryCtx.Done()`, then return `queryCtx.Err()`.

## Solution

```go
// QueryErrAfterRequestCancel builds the per-request context, derives the
// database query context from it, cancels the request (the client hung up),
// and returns the error the query context reports.
//
// Examples:
//
//	QueryErrAfterRequestCancel()                              => context.Canceled
//	errors.Is(QueryErrAfterRequestCancel(), context.Canceled) => true
//	the result is never nil
func QueryErrAfterRequestCancel() error {
	reqCtx, cancelReq := context.WithCancel(context.Background())
	queryCtx, cancelQuery := context.WithCancel(reqCtx)
	defer cancelQuery()

	cancelReq()

	<-queryCtx.Done()
	return queryCtx.Err()
}
```

## Walkthrough

- `cancelReq()` closes the request context's `Done()`.
- The runtime then cancels each registered child, closing the query context's `Done()` with the same reason.
- `queryCtx.Err()` therefore reports `context.Canceled`, and the database driver aborts the statement.

## Pitfalls

- Reading the child's `Err()` without waiting on its `Done()` can race with propagation — wait first.
- Skipping the child's `defer cancel()` still leaks it if the parent outlives the function; `go vet` flags it.
- Propagation is one-directional: down the tree only, never back up to the parent.
