# Rollout Phase Barrier

## Intuition

A barrier is a turnstile that opens only when the whole group has queued. The trap is reuse: if waiters looped on `arrived < need`, the reset to zero would look identical to "nobody has arrived yet" and they would sleep forever. Waiting on the phase number sidesteps that entirely.

## Approach

1. Clamp `n` and build the `Cond` over the mutex.
2. In `Wait`, take the lock and snapshot the current phase.
3. Increment `arrived`. If it reached `need`: reset it, bump the phase, `Broadcast`, return.
4. Otherwise loop `for phase == b.phase { cond.Wait() }`.

## Solution

```go
func NewBarrier(n int) *Barrier {
	if n < 1 {
		n = 1
	}
	b := &Barrier{need: n}
	b.cond = sync.NewCond(&b.mu)
	return b
}

func (b *Barrier) Wait() {
	b.mu.Lock()
	defer b.mu.Unlock()

	phase := b.phase
	b.arrived++
	if b.arrived == b.need {
		b.arrived = 0
		b.phase++
		b.cond.Broadcast()
		return
	}
	for phase == b.phase {
		b.cond.Wait()
	}
}
```

## Walkthrough

Eight regions, seven arrive: each snapshots phase 0, increments, and sleeps because the phase has not moved. The eighth increments to 8, resets `arrived` to 0, bumps the phase to 1, and broadcasts. The sleepers wake, see `0 != 1`, and leave. The counter is already zeroed, so phase two runs on the same barrier.

## Pitfalls

- Looping on `arrived < need`: after the reset the predicate is true again and everyone sleeps forever.
- `Signal` instead of `Broadcast` — one region proceeds, the rest hang.
- Reading the phase after incrementing, which can put a fast participant into the next round's predicate.
