# Batching a Stream

## Intuition

Batching converts a stream into a stream of groups. Two conditions end a
group: it filled up, or the stream ended. Handling only the first loses the
tail; handling only the second defeats the point.

## Approach

1. Guard `size <= 0` by draining and returning nil.
2. Range over `rows`, appending to the current batch; when it reaches `size`, store it and `make` a fresh one.
3. After the range, store the current batch if it is non-empty.

## Solution

```go
func BatchInserts(rows <-chan string, size int) [][]string {
	if size <= 0 {
		for range rows {
		}
		return nil
	}

	var batches [][]string
	batch := make([]string, 0, size)

	for row := range rows {
		batch = append(batch, row)
		if len(batch) == size {
			batches = append(batches, batch)
			batch = make([]string, 0, size)
		}
	}

	if len(batch) > 0 {
		batches = append(batches, batch)
	}
	return batches
}
```

## Walkthrough

With rows a, b, c and size 2: a and b fill the first batch, which is stored,
and a fresh batch is allocated. c goes into it, then the stream closes. The
partial batch `{c}` is non-empty, so it is flushed as the second batch.

## Pitfalls

- Dropping the final partial batch — the last rows of every load silently vanish.
- Reusing the same backing array via `batch = batch[:0]`: later appends overwrite the batch already stored.
- Appending an empty final batch when the row count divides evenly — check `len(batch) > 0` first.
