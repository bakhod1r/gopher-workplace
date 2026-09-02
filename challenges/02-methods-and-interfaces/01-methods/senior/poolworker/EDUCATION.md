# Worker Pool Methods

## Intuition

A pool turns "run this" into "queue this". The channel is the queue, the fixed
set of goroutines is the bound on concurrency, and the `WaitGroup` is the answer
to "is everything finished?". All three live inside the struct, so the caller
only sees `Start`, `Tasks` and `Wait`.

The shutdown protocol is worth stating out loud: the *producer* closes `Tasks`,
each worker's `range` then ends, `Done` fires, and `Wait` returns.

## Approach

1. Loop `p.Count` times.
2. Register the worker with the WaitGroup before starting it.
3. In the worker, range over the task channel and call each task.
4. Defer `Done` so the worker deregisters however it exits.

## Solution

```go
func (p *Pool) Start() {
	for i := 0; i < p.Count; i++ {
		p.wg.Add(1)
		go func() {
			defer p.wg.Done()
			for task := range p.Tasks {
				task()
			}
		}()
	}
}
```

## Walkthrough

Three workers start and block on an empty channel. The test queues five
closures, each incrementing a counter atomically, then closes `Tasks`. The
workers drain the queue between them — which worker runs which task is
unspecified — and once the channel is closed and empty every `range` ends,
every deferred `Done` runs, and `Wait` unblocks with the counter at 5.

## Pitfalls

- **`p.wg.Add(1)` inside the goroutine.** The scheduler may not have run it
  before `Wait` is called, so `Wait` sees a zero counter and returns while work
  is still pending.
- **Not calling `task()`.** The workers drain the channel and do nothing; the
  counter stays 0.
- **Closing `Tasks` inside the pool.** The producer owns the channel's lifetime;
  closing it from a consumer risks a send on a closed channel.
- **Copying the `Pool`.** It holds a `sync.WaitGroup`, which must not be copied —
  `go vet` flags it.

## Why `defer`

A task that panics would otherwise skip `Done` and deadlock `Wait` forever. The
deferred call keeps the accounting correct on every exit path.
