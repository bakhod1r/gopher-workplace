# Merge Runs Into The Caller's Buffer

## Intuition

A k-way merge is a repeated minimum over the run heads. Nothing needs copying and nothing needs sorting — the only state is where each run has got to, which is small enough to live on the stack.

## Approach

1. Sum the run lengths; return `dst` when the total is zero.
2. Keep cursors in a fixed local array when the run count allows, otherwise allocate.
3. Repeat `total` times: scan for the smallest head, append it, advance its cursor.

## Solution

```go
// Merge appends every element of the sorted runs to dst in ascending
// order and returns the extended slice.
//
// With room in dst, the merge must allocate nothing: the cursors are the
// only state it needs.
//
// Examples:
//
// 	Merge(nil, [][]int{{1, 3}, {2}}) => []int{1, 2, 3}
func Merge(dst []int, runs [][]int) []int {
	total := 0
	for _, r := range runs {
		total += len(r)
	}
	if total == 0 {
		return dst
	}
	var cursors [16]int
	pos := cursors[:0]
	if len(runs) <= cap(pos) {
		pos = cursors[:len(runs)]
	} else {
		pos = make([]int, len(runs))
	}
	for n := 0; n < total; n++ {
		best := -1
		for i, r := range runs {
			if pos[i] >= len(r) {
				continue
			}
			if best < 0 || r[pos[i]] < runs[best][pos[best]] {
				best = i
			}
		}
		dst = append(dst, runs[best][pos[best]])
		pos[best]++
	}
	return dst
}
```

## Walkthrough

Merging {1,4,7}, {2,5,8} and {3,6,9} advances one cursor per step. The cursor array lives in the frame, so with room in `dst` the whole merge allocates nothing.

## Pitfalls

- Allocating the cursor slice unconditionally, which is one allocation per call.
- Skipping exhausted runs incorrectly and indexing past a run's end.
- A heap is the right structure for many runs; the linear scan is fine for a few.
