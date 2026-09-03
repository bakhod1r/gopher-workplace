# Log Shipper Worker Pool

## Intuition

A worker pool inverts the fan-out: goroutines are created once and jobs flow to them. The channel close is the entire shutdown protocol — no quit flag, no counter, no timeout. Each `range` loop ends on its own when the work runs out, and the WaitGroup tells the parent when the last one has.

## Approach

1. Allocate `errs := make([]error, len(lines))` and return early for an empty input.
2. Clamp `workers` into `[1, len(lines)]`.
3. Create `jobs := make(chan int)` and start `workers` goroutines, each ranging over `jobs` and writing `ship(lines[i])` into `errs[i]`.
4. Send every index from the parent, then `close(jobs)`.
5. `wg.Wait()` and return `errs`.

## Solution

```go
// Package logworkerpool — Gopher Workplace challenge.
package logworkerpool

import "sync"

// ShipLines ships every log line through a fixed pool of worker goroutines and
// returns one slot per line, in input order: nil when the line was accepted,
// the shipper's error when it was not. The pool size is fixed regardless of how
// many lines arrive, so a burst of a million lines still costs workers
// goroutines and workers open connections.
//
// A workers value of zero or less is treated as one.
//
// Examples:
//
//	ShipLines([]string{"ok", "bad"}, 2, ship)  => [<nil> errRejected]
//	ShipLines([]string{"ok"}, 8, ship)         => [<nil>]
//	ShipLines(nil, 4, ship)                    => []
func ShipLines(lines []string, workers int, ship func(line string) error) []error {
	errs := make([]error, len(lines))
	if len(lines) == 0 {
		return errs
	}
	if workers <= 0 {
		workers = 1
	}
	if workers > len(lines) {
		workers = len(lines)
	}

	jobs := make(chan int)

	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				errs[i] = ship(lines[i])
			}
		}()
	}

	for i := range lines {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
	return errs
}
```

## Walkthrough

- With six lines and three workers the shipper's peak counter never exceeds three, yet all six lines are shipped.
- `more_workers_than_lines` clamps the pool to two so no worker is started with nothing to do.
- `non_positive_workers` clamps up to one and the pool degenerates to a sequential shipper — still correct.
- Because each worker writes only `errs[i]` for the index it pulled, no two workers ever touch the same slot.

## Pitfalls

- Forgetting `close(jobs)`: every worker blocks forever in `range` and `wg.Wait()` deadlocks.
- Closing the channel from a worker — a second close panics, and the remaining sends panic too.
- Sending the line value instead of its index, which leaves the worker unable to find its result slot.
- Starting one goroutine per line and calling it a pool; the connection budget is then unbounded again.
