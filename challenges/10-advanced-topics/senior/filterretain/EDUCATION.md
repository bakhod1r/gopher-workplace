# The Filter That Keeps The Whole Batch Alive

## Intuition

In-place compaction is the right move when the input dies with the result. Here the result outlives the batch, so a view of the batch is a leak of the batch — the survivors have to move to storage of their own.

## Approach

1. First pass: count the survivors.
2. Allocate a result with exactly that capacity.
3. Second pass: append the survivors.

## Solution

```go
// Record is one ingested item.
type Record struct {
	ID   int
	Size int
	Pad  [64]byte
}

// Keep returns the records whose Size is at least min.
//
// Typical batches are huge and typical results are tiny. The result must
// not keep the batch's storage alive once the caller drops the batch.
//
// Examples:
//
// 	Keep(batch, 100) => only the large records
func Keep(records []Record, min int) []Record {
	n := 0
	for _, r := range records {
		if r.Size >= min {
			n++
		}
	}
	out := make([]Record, 0, n)
	for _, r := range records {
		if r.Size >= min {
			out = append(out, r)
		}
	}
	return out
}
```

## Walkthrough

16384 records of 80 bytes is about 1.3 MiB. Returning `records[:3]` keeps every byte of it reachable; copying three records into a 3-element array keeps 240 bytes.

## Pitfalls

- `records[:k:k]` — caps the capacity, still points at the batch.
- Appending to a nil slice, which is correct but allocates through several growth steps.
