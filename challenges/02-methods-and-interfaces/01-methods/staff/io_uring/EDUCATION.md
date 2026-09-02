# Submission Ring

## Intuition

A ring buffer is a queue that never allocates. Two cursors chase each other
around a fixed array; the only interesting question is how to tell a full ring
from an empty one, since in both cases `head == tail`.

Two classic answers exist: keep a slot permanently empty, or keep an explicit
count. This puzzle uses the count, because it makes both guards a single
comparison and lets the ring use all `ringSize` slots.

## Approach

1. Guard first — full for `Submit`, empty for `Complete`.
2. Access the array at the relevant cursor.
3. Advance that cursor modulo the capacity.
4. Adjust `count`.

## Solution

```go
func (r *Ring) Submit(op int) bool {
	if r.count == ringSize {
		return false
	}
	r.buf[r.tail] = op
	r.tail = (r.tail + 1) % ringSize
	r.count++
	return true
}

func (r *Ring) Complete() (int, bool) {
	if r.count == 0 {
		return 0, false
	}
	op := r.buf[r.head]
	r.head = (r.head + 1) % ringSize
	r.count--
	return op, true
}
```

## Walkthrough

The wrap-around test is the one that matters. With `ringSize == 4`:

| step | head | tail | count |
|------|------|------|-------|
| submit 0,1,2,3 | 0 | 0 | 4 |
| complete ×3 | 3 | 0 | 1 |
| submit 100,101 | 3 | 2 | 3 |

`tail` has wrapped past the end and is now *behind* `head`, reusing slots 0 and
1 — which the earlier completions freed. Draining yields `3, 100, 101`: exactly
submission order, with the reuse invisible from outside.

## Pitfalls

- **Advancing before the guard.** A rejected `Submit` that already moved `tail`
  desynchronizes the cursors permanently.
- **Using `head == tail` as the full/empty test.** Ambiguous without the count;
  this is the bug that makes a ring silently drop every entry once it fills.
- **`r.tail++` without the modulo.** Indexes out of range on the fifth submit.
- **Value receiver.** `[ringSize]int` is an array, so the whole struct including
  the buffer is copied and every write is thrown away.

## Why the real thing is harder

Actual `io_uring` shares this structure with the kernel across mmap'd memory, so
the cursor updates need atomic stores with release/acquire ordering rather than
plain increments. The data-structure logic is the same; the memory model is
where the difficulty moves.
