# Asynchronous Methods

## Intuition

Methods that start background work often return a signal channel. The caller can
decide whether to block `<-f.FetchAsync()` or do other work.

## Solution

```go
func (f *Fetcher) FetchAsync(id string) <-chan struct{} {
	done := make(chan struct{})
	go func() {
		f.Fetch(id)
		close(done)
	}()
	return done
}
```

## Walkthrough

- `done` channel created.
- Goroutine starts.
- Main thread returns `done`.
- Goroutine completes `Fetch`, closes `done`.
- Caller's `<-done` unblocks.
