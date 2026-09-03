# Amortising A Fixed Cost

## Intuition

Buffer until you have enough to make the expensive call worthwhile, then make it once. The buffer is reused, so the only cost per item is an append.

## Approach

1. `Add` appends, then flushes if the buffer reached the size.
2. A private `flush` handles the nil callback, the counter, and the reset.
3. `Close` flushes only when something is buffered.

## Solution

```go
func (b *Batcher) size() int {
	if b.Size <= 0 {
		return 1
	}
	return b.Size
}

func (b *Batcher) flush() {
	if len(b.buf) == 0 {
		return
	}
	if b.Flush != nil {
		b.Flush(b.buf)
	}
	b.sent++
	b.buf = b.buf[:0]
}

func (b *Batcher) Add(v int) {
	b.buf = append(b.buf, v)
	if len(b.buf) >= b.size() {
		b.flush()
	}
}

func (b *Batcher) Close() { b.flush() }

func (b *Batcher) Flushes() int { return b.sent }
```

## Walkthrough

Putting the empty check inside `flush` is what makes `Close` idempotent: the second call finds an empty buffer and returns without invoking the callback.

## Pitfalls

- `b.buf = nil` after flushing, which reallocates the batch on every cycle.
- Passing the live buffer to a callback that stores it — the next batch overwrites it.
- Forgetting `Close`, so the last partial batch is silently lost.
