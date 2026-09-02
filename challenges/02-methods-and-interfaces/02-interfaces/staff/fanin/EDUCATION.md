# Fan-In With Shutdown

## Intuition

The leak lives in the send, not the receive. A forwarder blocked on `out <- v` with nobody reading will sit there forever, so the send itself has to be a `select` that can hear the shutdown signal.

## Approach

1. Start one goroutine per input, each holding a `wg` ticket.
2. The receive is a `select` over `in` and `done`.
3. The send is another `select` over `out <- v` and `done`.
4. A separate goroutine waits on the WaitGroup and closes `out` exactly once.

## Solution

```go
func (Fan) Merge(done <-chan struct{}, ins ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	for _, in := range ins {
		wg.Add(1)
		go func(in <-chan int) {
			defer wg.Done()
			for {
				select {
				case v, ok := <-in:
					if !ok {
						return
					}
					select {
					case out <- v:
					case <-done:
						return
					}
				case <-done:
					return
				}
			}
		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
```

## Walkthrough

`TestEarlyExitDoesNotLeak` reads one value from three infinite producers and closes `done`. Only the nested `select` on the send lets those forwarders escape.

## Pitfalls

- A plain `out <- v` — the forwarder blocks forever once the consumer stops reading.
- Closing `out` from each forwarder, which panics on the second close.
- Capturing the loop variable instead of passing `in` as an argument, so every goroutine reads the same channel.
