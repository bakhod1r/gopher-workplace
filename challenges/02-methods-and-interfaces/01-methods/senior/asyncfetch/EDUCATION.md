# Asynchronous Methods

## Intuition

An async API separates *starting* work from *waiting* for it. The method starts
a goroutine and returns a handle immediately; the handle here is a channel whose
closing means "finished". `struct{}` is the payload precisely because there is
no payload — only the event matters, and an empty struct costs zero bytes.

## Approach

1. Create the signal channel.
2. Launch the goroutine, doing the real work inside it.
3. Close the channel when that work returns.
4. Return the channel to the caller straight away.

## Solution

```go
func (f *Fetcher) FetchAsync(id string) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		defer close(done)
		f.Fetch(id)
	}()
	return done
}
```

## Walkthrough

`FetchAsync` returns while the goroutine is still running, so the test's
`select` blocks on `<-done`. `Fetch` takes the mutex, writes `Result`, and
releases it; the deferred `close` then fires and the receive completes. Because
the close happens after the write, the caller is guaranteed to observe
`"data: abc"` — the channel provides the happens-before edge.

The `time.After(time.Second)` arm exists to fail loudly rather than hang if the
channel is never closed.

## Pitfalls

- **Closing before the work.** `close(done)` before `f.Fetch(id)` releases the
  caller too early; it reads `Result` while the goroutine is still writing it —
  a data race, and usually an empty string.
- **Never closing.** The caller blocks forever; the test's timeout is what turns
  that into a readable failure instead of a hang.
- **Calling `Fetch` synchronously and then returning a closed channel.** It
  passes the test, but the method is no longer asynchronous.
- **Returning `chan struct{}` instead of `<-chan struct{}`.** The caller could
  then close it, and a second close panics.

## Why `defer close(done)`

If `Fetch` ever panics, the deferred close still runs, so the waiter is released
instead of deadlocking. That habit matters more as the goroutine body grows.
