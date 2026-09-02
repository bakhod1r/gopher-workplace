# Semaphore

## Intuition

A buffered channel *is* a counting semaphore: capacity is the permit count, a send takes one, a receive returns one. No mutex, no counter to get wrong.

## Approach

1. `Acquire` sends into `slots`, blocking when full.
2. `TryAcquire` wraps the same send in a `select` with `default`.
3. `Release` receives from `slots`.
4. `RunLimited` acquires *before* starting each goroutine, then releases and marks done via `defer`.

## Solution

```go
func (s *Semaphore) Acquire() { s.slots <- struct{}{} }

func (s *Semaphore) TryAcquire() bool {
	select {
	case s.slots <- struct{}{}:
		return true
	default:
		return false
	}
}

func (s *Semaphore) Release() { <-s.slots }

func RunLimited(jobs []Job, limit int) {
	if limit < 1 {
		limit = 1
	}
	sem := NewSemaphore(limit)

	var wg sync.WaitGroup
	for _, j := range jobs {
		wg.Add(1)
		sem.Acquire()
		go func(j Job) {
			defer wg.Done()
			defer sem.Release()
			j.Do()
		}(j)
	}
	wg.Wait()
}
```

## Walkthrough

Acquiring in the *loop* rather than inside the goroutine is what bounds the goroutine count too: the loop blocks when the semaphore is full, so at most `limit` goroutines exist at once.

## Pitfalls

- Acquiring inside the goroutine, which still spawns one goroutine per job — the concurrency is bounded but the memory is not.
- `defer sem.Release()` missing on an early return path, which permanently loses a slot.
- Releasing more than you acquire — the receive blocks and the pipeline stalls.
