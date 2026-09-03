# A Permit To Proceed

## Intuition

Fill a buffered channel with tokens and taking one is receiving; giving it back is sending. Or leave it empty and treat a send as taking a slot — either convention works, as long as everything agrees.

## Approach

1. `New` builds a channel with capacity `n`; an occupied slot means a permit is taken.
2. `Acquire` sends, `Release` receives non-blockingly, `TryAcquire` sends with a `default`.
3. `Available` is `cap - len`.

## Solution

```go
func New(n int) *Sema {
	if n < 1 {
		n = 1
	}
	return &Sema{slots: make(chan struct{}, n)}
}

func (s *Sema) Acquire() { s.slots <- struct{}{} }

func (s *Sema) TryAcquire() bool {
	select {
	case s.slots <- struct{}{}:
		return true
	default:
		return false
	}
}

func (s *Sema) Release() {
	select {
	case <-s.slots:
	default:
	}
}

func (s *Sema) Available() int { return cap(s.slots) - len(s.slots) }
```

## Walkthrough

The `default` in `Release` is what makes an unmatched release harmless: with no token to take back, the receive would block forever, so the operation is simply dropped.

## Pitfalls

- A blocking `Release`, which deadlocks the first time a caller releases a permit twice.
- `defer s.Release()` after a `TryAcquire` that returned false, handing back a permit you never took.
- Sizing the gate by intuition; the right limit comes from what the protected resource can actually sustain.
