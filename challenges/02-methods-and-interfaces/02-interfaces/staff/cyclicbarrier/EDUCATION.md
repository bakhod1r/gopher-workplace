# Cyclic Barrier

## Intuition

A one-shot barrier is easy; reuse is where it gets subtle. The round counter is what distinguishes "my group has been released" from "a later group is filling up", so a woken goroutine can tell whether the wait is really over.

## Approach

1. Record `myRound` on arrival and increment `count`.
2. The last arrival resets `count`, increments `round`, and broadcasts.
3. Everyone else waits while `b.round == myRound`.
4. All parties return the round index they completed.

## Solution

```go
func (b *Barrier) Await() int {
	b.mu.Lock()
	defer b.mu.Unlock()

	myRound := b.round
	b.count++

	if b.count == b.Parties {
		b.count = 0
		b.round++
		b.cond.Broadcast()
		return myRound
	}

	for b.round == myRound {
		b.cond.Wait()
	}
	return myRound
}
```

## Walkthrough

`TestNoOneRunsAhead` has four parties run fifty rounds and asserts each goroutine sees round indexes 0..49 in order. Waiting on `count == 0` instead of the round would let a fast goroutine from round N+1 be mistaken for a straggler.

## Pitfalls

- Waiting on `count` rather than the round: a new arrival for the next round changes `count` and wakes the wrong goroutines.
- `Signal` instead of `Broadcast`, releasing one party per round.
- Resetting `count` outside the lock, which loses arrivals under contention.
