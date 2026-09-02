# Goroutine Leak

## Intuition

A goroutine blocked on a channel nobody will ever write to is invisible: no error, no crash, just memory that never comes back. Every goroutine needs an owner who can stop it and a `select` case that lets it hear about it.

## Approach

1. `Start` runs a loop with `defer close(w.done)` so exit is always signalled.
2. The `select` handles both a value from `in` (including closure) and the `stop` signal.
3. `Stop` closes `stop` inside a `sync.Once`, then waits on `done`.
4. `Wait` just blocks on `done`, for callers that only care about the input closing.

## Solution

```go
func (w *Watcher) Start(in <-chan int, s Sink) {
	go func() {
		defer close(w.done)
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				s.Observe(v)
			case <-w.stop:
				return
			}
		}
	}()
}

func (w *Watcher) Stop() {
	w.stopOnce.Do(func() {
		close(w.stop)
	})
	<-w.done
}
```

## Walkthrough

`TestStopWhileBlockedOnSend` is the leak in miniature: with only `case v := <-in`, the goroutine parks forever on a channel nobody writes. The `stop` case is the escape hatch, and `NumGoroutine` proves it worked.

## Pitfalls

- `for v := range in` with no stop case — the goroutine lives as long as the channel does.
- `close(w.stop)` without `sync.Once`, so a second `Stop` panics.
- Signalling stop but never waiting on `done`, so the caller cannot tell when the goroutine actually exited.
