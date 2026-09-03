# Keys That Were Overwritten By The Next Batch

## Intuition

`unsafe.String` is only sound over bytes that will never change. An arena is the opposite of that. The fix is not to abandon the technique but to give it memory that qualifies — one block per batch, written once.

## Approach

1. Sum the valid spans' lengths for the block's capacity.
2. Append each valid span's bytes into the block, remembering where it started.
3. Wrap `block[start:]` with `unsafe.String` and the span's length.

## Solution

```go
import "unsafe"

// Intern returns one string per span of arena, for the caller to keep.
//
// arena is a scratch region the caller refills between batches, so the
// strings must own their bytes. Copying them into one block keeps the cost
// to a single allocation for the whole batch.
//
// Examples:
//
// 	Intern([]byte("abcd"), [][2]int{{0,2},{2,4}}) => []string{"ab", "cd"}
func Intern(arena []byte, spans [][2]int) []string {
	if len(spans) == 0 {
		return nil
	}
	total := 0
	for _, sp := range spans {
		lo, hi := sp[0], sp[1]
		if lo < 0 || hi > len(arena) || lo >= hi {
			continue
		}
		total += hi - lo
	}
	block := make([]byte, 0, total)
	out := make([]string, 0, len(spans))
	for _, sp := range spans {
		lo, hi := sp[0], sp[1]
		if lo < 0 || hi > len(arena) || lo >= hi {
			out = append(out, "")
			continue
		}
		start := len(block)
		block = append(block, arena[lo:hi]...)
		out = append(out, unsafe.String(unsafe.SliceData(block[start:]), hi-lo))
	}
	return out
}
```

## Walkthrough

Before the fix, twenty-six batches insert twenty-six keys that all point into a four-byte arena, so the map ends up full of entries whose keys now read "zzzz". After it, each batch's bytes are copied once and the keys never change.

## Pitfalls

- Appending to `block` after taking a string over it — a reallocation would leave the earlier strings pointing at the old array. Sizing the block up front is what makes this safe.
- Allocating per span, which is correct and defeats the purpose.
