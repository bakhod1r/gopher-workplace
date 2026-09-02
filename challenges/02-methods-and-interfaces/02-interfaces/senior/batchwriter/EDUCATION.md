# Batching Writer

## Intuition

Batching converts N expensive calls into N/size of them. The subtlety is ownership: the sink is handed the live buffer, so it must copy anything it intends to keep.

## Approach

1. `Write` appends, then flushes when `len(buf) >= size`.
2. `Flush` returns early when empty, otherwise hands the buffer to the sink.
3. Reset with `b.buf = b.buf[:0]` so the backing array is reused.
4. `Buffered` reports the pending count.

## Solution

```go
func (b *BatchWriter) Write(record string) {
	b.buf = append(b.buf, record)
	if len(b.buf) >= b.size {
		b.Flush()
	}
}

func (b *BatchWriter) Flush() {
	if len(b.buf) == 0 {
		return
	}
	b.sink.WriteBatch(b.buf)
	b.buf = b.buf[:0]
}

func (b *BatchWriter) Buffered() int { return len(b.buf) }
```

## Walkthrough

`RecordingSink.WriteBatch` copies before storing — that is the contract this design forces on sinks. `TestBufferIsReused` then proves the writer never reallocates its buffer.

## Pitfalls

- `b.buf = nil` after flushing, which allocates a new array on the next write.
- Storing the passed slice in the sink without copying: the next batch overwrites it.
- Forgetting the final `Flush`, silently dropping the last partial batch.
