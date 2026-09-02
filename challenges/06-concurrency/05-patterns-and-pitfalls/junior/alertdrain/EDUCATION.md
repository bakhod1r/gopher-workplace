# Done Channels and select

## Intuition

Closing a channel wakes *every* goroutine blocked on it, which makes
`close(done)` a one-to-many cancellation signal. Pairing it with the work
channel in a `select` means a consumer can never get stuck on a dead producer
during shutdown.

## Approach

1. Loop forever with a `select`.
2. `case a, ok := <-alerts`: if `!ok` the stream ended — return the count; otherwise increment.
3. `case <-done`: return the count immediately.

## Solution

```go
func CountAlerts(done <-chan struct{}, alerts <-chan string) int {
	count := 0
	for {
		select {
		case _, ok := <-alerts:
			if !ok {
				return count
			}
			count++
		case <-done:
			return count
		}
	}
}
```

## Walkthrough

With `done` open and three buffered alerts then a close: three iterations take
the alerts case and increment; the fourth sees `ok == false` and returns 3.
With `done` already closed and no alerts, the first select takes the done case
and returns 0.

## Pitfalls

- Sending a value on `done` instead of closing it — only one of the waiting goroutines would wake up.
- Ignoring `ok`: a closed channel yields zero values forever, so the loop would spin at 100% CPU.
- Checking `done` outside the select, where a blocking receive on `alerts` can never reach it.
