# MPMC Ring Buffer

## Intuition

The sequence number on each slot is the synchronisation. `seq == pos` means "a producer's turn"; `seq == pos+1` means "a consumer's turn". Publishing the new sequence is the release that hands the slot to the other side.

## Approach

1. `Enqueue` loads the tail, inspects the target slot's sequence, and CASes the tail when the slot is free.
2. After winning the slot, write the value and store `pos+1` to publish it.
3. `Dequeue` mirrors this: it needs `seq == pos+1`, then republishes `pos + mask + 1` to mark the slot free for the next lap.
4. A sequence behind the expected value means full (enqueue) or empty (dequeue).

## Solution

```go
func (r *Ring) Enqueue(v int) bool {
	for {
		pos := r.tail.Load()
		s := &r.buf[pos&r.mask]
		seq := s.seq.Load()

		switch {
		case seq == pos:
			if r.tail.CompareAndSwap(pos, pos+1) {
				s.val = v
				s.seq.Store(pos + 1)
				return true
			}
		case seq < pos:
			return false // full
		}
	}
}

func (r *Ring) Dequeue() (int, bool) {
	for {
		pos := r.head.Load()
		s := &r.buf[pos&r.mask]
		seq := s.seq.Load()

		switch {
		case seq == pos+1:
			if r.head.CompareAndSwap(pos, pos+1) {
				v := s.val
				s.seq.Store(pos + r.mask + 1)
				return v, true
			}
		case seq < pos+1:
			return 0, false // empty
		}
	}
}
```

## Walkthrough

Storing `pos + mask + 1` on dequeue is what makes the slot available exactly one lap later. Storing `pos+1` there instead would let a producer overwrite a slot a consumer had not finished with.

## Pitfalls

- Deriving fullness from `tail - head` without the per-slot sequence — a producer can then write a slot a consumer is still reading.
- Writing `s.val` before winning the CAS, which lets two producers race on one slot.
- Using a non-power-of-two capacity and masking anyway, which corrupts the index mapping.
