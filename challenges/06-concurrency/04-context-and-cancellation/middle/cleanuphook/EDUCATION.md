# Connection Cleanup Hook

## Intuition

`AfterFunc` is a callback on cancellation with a cancel of its own. That second cancel is what makes early release expressible: whoever wins gets to name the reason, and the loser's answer is simply `false`.

## Approach

1. Build the Hub with a `closed` channel; register the hook with `AfterFunc` and store the stop func.
2. `finish` uses a `sync.Once` to record who tore down and close the channel.
3. `Release` calls the stop func; on `true` it finishes as `"release"`, and it returns that result.
4. `Wait` receives from `closed`, then returns the recorded reason.

## Solution

```go
import (
	"context"
	"sync"
)

type Hub struct {
	closed chan struct{}
	stop   func() bool
	once   sync.Once
	by     string
	mu     sync.Mutex
}

func Register(ctx context.Context) *Hub {
	h := &Hub{closed: make(chan struct{})}
	h.stop = context.AfterFunc(ctx, func() {
		h.finish("context")
	})
	return h
}

func (h *Hub) finish(by string) {
	h.once.Do(func() {
		h.mu.Lock()
		h.by = by
		h.mu.Unlock()
		close(h.closed)
	})
}

func (h *Hub) Release() bool {
	stopped := h.stop()
	if stopped {
		h.finish("release")
	}
	return stopped
}

func (h *Hub) Wait() string {
	<-h.closed
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.by
}
```

## Walkthrough

A client closes cleanly: `Release` calls the stop func, which returns `true` because the hook has not fired, so the Hub records `"release"` and closes. If instead the request is cancelled first, `AfterFunc` runs the hook, which records `"context"`; the later `Release` gets `false` from the stop func and changes nothing.

## Pitfalls

- Finishing as `"release"` even when the stop func returned `false` — the reason then contradicts what really happened.
- Closing the channel on both paths without a `sync.Once`: closing a closed channel panics.
- Replacing `AfterFunc` with a goroutine on `<-ctx.Done()` that has no way to be stopped.
