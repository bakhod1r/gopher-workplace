# Abort an Upload

## Intuition

`WithCancel` hands back two things: a child context and a switch. Flipping the switch closes the child's `Done()` channel and stamps `ctx.Err()` with `context.Canceled` — permanently. Reading `Err()` is how the storage client, and later the audit log, learns the transfer was aborted rather than timed out.

## Approach

1. Derive a cancellable context from `context.Background()`.
2. Call `cancel()`.
3. Return `ctx.Err()`.

## Solution

```go
// AbortUpload starts a cancellable upload context, aborts it because the user
// pressed Cancel in the UI, and returns the reason recorded on the context.
//
// Examples:
//
//	AbortUpload()                                     => context.Canceled
//	errors.Is(AbortUpload(), context.Canceled)        => true
//	errors.Is(AbortUpload(), context.DeadlineExceeded) => false
func AbortUpload() error {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx.Err()
}
```

## Walkthrough

- Before `cancel()`, `ctx.Err()` is nil — the upload is still live.
- `cancel()` closes `ctx.Done()` and stores `context.Canceled`.
- Cancellation is sticky: extra `cancel()` calls change nothing and `Err()` keeps returning the same sentinel.

## Pitfalls

- Reading `ctx.Err()` *before* calling `cancel()` yields nil.
- `defer cancel()` would run after the read, so the observed error would be nil — call it inline here.
- `context.Canceled` is a value, not a type; compare with `==` or `errors.Is`, never by matching the message string.
