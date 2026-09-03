# Bound The Work In Flight

## Intuition

A buffered channel already implements counting with blocking — sending is acquiring and the capacity is the limit. The only thing it lacks is a way out, and `select` supplies that.

## Approach

1. `select` on sending an empty struct into `s.slots` and on `<-done`.
2. Return true for the first, false for the second.

## Solution

```go
// Sem is a counting semaphore of fixed capacity.
type Sem struct {
	slots chan struct{}
}

// NewSem returns a semaphore permitting n concurrent holders.
func NewSem(n int) *Sem {
	if n < 1 {
		n = 1
	}
	return &Sem{slots: make(chan struct{}, n)}
}

// Release returns one slot. It must only be called after a successful Acquire.
func (s *Sem) Release() {
	select {
	case <-s.slots:
	default:
	}
}

// Held reports how many slots are currently taken.
func (s *Sem) Held() int { return len(s.slots) }

// Acquire takes one slot, blocking until one is free, and reports whether
// it got one.
//
// It must also give up when done is closed, so a cancelled caller does not
// wait forever on a saturated semaphore.
//
// Examples:
//
// 	s := NewSem(2); s.Acquire(done) => true twice, then blocks
func (s *Sem) Acquire(done <-chan struct{}) bool {
	select {
	case s.slots <- struct{}{}:
		return true
	case <-done:
		return false
	}
}
```

## Walkthrough

With 32 workers and a limit of 4, at most four sends fit in the buffer; the rest block in the `select` until a `Release` receives one out, or `done` closes.

## Pitfalls

- Checking `len(s.slots) < cap(s.slots)` first — the state can change before the send.
- Returning true on the `done` branch, so the caller releases a slot it never took.
- Releasing without a matching acquire, which the `default` in `Release` tolerates but which breaks the count.
