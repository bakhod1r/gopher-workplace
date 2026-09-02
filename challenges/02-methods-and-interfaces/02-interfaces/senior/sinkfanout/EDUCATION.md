# Fan-Out Sink

## Intuition

Fan-out is embarrassingly parallel, and the only shared state is the failure tally. Give each goroutine its own slot and the race disappears without a mutex on the hot path.

## Approach

1. `MemSink.Write` appends under its mutex, since many goroutines may share one sink.
2. `ErrSink.Write` returns the sentinel.
3. `FanOut` allocates `failed` with one slot per sink.
4. Each goroutine writes only its own index; after `wg.Wait()`, count the flags.

## Solution

```go
func (m *MemSink) Write(event string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Events = append(m.Events, event)
	return nil
}

func (ErrSink) Write(event string) error { return ErrSinkFailed }

func FanOut(sinks []Sink, event string) int {
	failed := make([]bool, len(sinks))

	var wg sync.WaitGroup
	for i, s := range sinks {
		wg.Add(1)
		go func(i int, s Sink) {
			defer wg.Done()
			if err := s.Write(event); err != nil {
				failed[i] = true
			}
		}(i, s)
	}
	wg.Wait()

	n := 0
	for _, f := range failed {
		if f {
			n++
		}
	}
	return n
}
```

## Walkthrough

`TestFanOutRepeated` sends the same `MemSink` 50 events while a sibling sink fails every time; under `-race` the mutex in `Write` and the per-index writes in `FanOut` keep it clean.

## Pitfalls

- `n++` inside the goroutines — a data race and a wrong count.
- Returning early on the first failure, so later sinks never receive the event.
- Sharing a `MemSink` without its mutex: concurrent `append` corrupts the slice.
