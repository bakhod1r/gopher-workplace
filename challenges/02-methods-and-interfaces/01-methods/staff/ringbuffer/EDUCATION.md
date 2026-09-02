# Ring Buffer

## Intuition

A ring buffer is a queue with a hard memory ceiling: it never grows, never
allocates after construction, and rejects work instead of consuming unbounded
memory. That property is why it shows up in audio pipelines, network drivers and
lock-free queues — anywhere a slow consumer must not be able to exhaust RAM.

The classic implementation question is how to distinguish full from empty when
both leave `head == tail`. Keeping an explicit `size` is the simplest answer and
lets the buffer use every slot.

## Approach

1. Guard on `size` first.
2. Access the array at the relevant cursor.
3. Advance that cursor modulo the capacity.
4. Adjust `size`.

## Solution

```go
func (r *RingBuffer) Push(val int) error {
	if r.size == len(r.data) {
		return ErrFull
	}
	r.data[r.tail] = val
	r.tail = (r.tail + 1) % len(r.data)
	r.size++
	return nil
}

func (r *RingBuffer) Pop() (int, error) {
	if r.size == 0 {
		return 0, ErrEmpty
	}
	val := r.data[r.head]
	r.head = (r.head + 1) % len(r.data)
	r.size--
	return val, nil
}
```

## Walkthrough

The wrap test, with capacity 3:

| step | head | tail | size | array |
|------|------|------|------|-------|
| push 1,2,3 | 0 | 0 | 3 | [1 2 3] |
| pop ×2 | 2 | 0 | 1 | [1 2 3] |
| push 4,5 | 2 | 2 | 3 | [4 5 3] |

Slots 0 and 1 now hold the new values while `head` still points at the surviving
`3`. Draining yields `3, 4, 5` — FIFO order preserved across the wrap, with the
stale copies of 1 and 2 simply overwritten.

## Pitfalls

- **Mutating before the guard.** A rejected `Push` that already moved `tail`
  corrupts the queue permanently; the `Len()` assertion after `ErrFull` is what
  catches it.
- **Using `head == tail` as the fullness test.** Ambiguous — the buffer either
  refuses everything or overwrites live data.
- **`errors.New` inside the method.** Each call creates a distinct error value,
  so `errors.Is(err, ErrFull)` is false and callers cannot classify the failure.
- **Missing modulo.** Index out of range on the first wrap.

## Sentinel errors vs error types

A package-level `var ErrFull = errors.New(...)` is the Go idiom for a failure
the caller must be able to *test* for, like `io.EOF` or `sql.ErrNoRows`.
`errors.Is` walks the wrap chain, so callers still recognise it after a
`fmt.Errorf("%w", ...)` several layers up.
