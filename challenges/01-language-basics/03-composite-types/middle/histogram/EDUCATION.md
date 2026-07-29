# Bucketing into a histogram

## Intuition

The bin for a value is `v / size` (integer division). Size the result to hold the
largest bin, then count:

```go
maxBin := 0
for _, v := range xs { if v/size > maxBin { maxBin = v / size } }
bins := make([]int, maxBin+1)
for _, v := range xs { bins[v/size]++ }
```

## Approach

1. Empty input -> empty slice.
2. Compute max bin = max(v/size).
3. Allocate counts of length maxBin+1.
4. Increment counts[v/size] for each v.
5. Return counts.

## Solution

```go
func Bucket(xs []int, size int) []int {
	if len(xs) == 0 {
		return []int{}
	}
	maxBin := 0
	for _, v := range xs {
		if b := v / size; b > maxBin {
			maxBin = b
		}
	}
	out := make([]int, maxBin+1)
	for _, v := range xs {
		out[v/size]++
	}
	return out
}
```

## Walkthrough

values/10: 5->0,12->1,15->1,25->2,3->0. maxBin=2, len 3. counts bin0=2,bin1=2,bin2=1 -> [2,2,1].

## Pitfalls

- Guard `size > 0` to avoid divide-by-zero.
- Allocate `maxBin+1` bins, or the highest value indexes out of range.
- This assumes non-negative values; negatives need an offset.
