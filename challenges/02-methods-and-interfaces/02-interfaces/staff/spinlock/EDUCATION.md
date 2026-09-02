# Spin Lock

## Intuition

A spin lock trades CPU for latency: no scheduler involvement, but a waiter burns a core. The bounded yield is what keeps it from being catastrophic — without it, a spinner on the same P as the holder can prevent the holder from ever running.

## Approach

1. `TryLock` is a single `CompareAndSwap(false, true)`.
2. `Lock` loops on that CAS, counting spins and calling `runtime.Gosched` past the threshold.
3. `Unlock` CASes `true -> false` and panics if it was not held.
4. `Locked` is a plain atomic load.

## Solution

```go
func (l *SpinLock) Lock() {
	spins := 0
	for !l.held.CompareAndSwap(false, true) {
		spins++
		if spins >= spinsBeforeYield {
			runtime.Gosched()
			spins = 0
		}
	}
}

func (l *SpinLock) TryLock() bool {
	return l.held.CompareAndSwap(false, true)
}

func (l *SpinLock) Unlock() {
	if !l.held.CompareAndSwap(true, false) {
		panic("spinlock: unlock of an unlocked lock")
	}
}
```

## Walkthrough

`TestMutualExclusion` increments a plain `int` from 1000 goroutines. If the lock is wrong, `-race` reports it and the count comes out short — the test is the exclusion proof.

## Pitfalls

- `l.held.Store(true)` after a `Load()` check — the check and the set are separate, so two goroutines both enter.
- Spinning without ever yielding, which can livelock on a single-P runtime.
- Using a spin lock for anything but a very short critical section: for real work, `sync.Mutex` is faster and fairer.
