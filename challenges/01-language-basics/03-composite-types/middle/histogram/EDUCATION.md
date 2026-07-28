# Bucketing into a histogram

## The idea

The bin for a value is `v / size` (integer division). Size the result to hold the
largest bin, then count:

```go
maxBin := 0
for _, v := range xs { if v/size > maxBin { maxBin = v / size } }
bins := make([]int, maxBin+1)
for _, v := range xs { bins[v/size]++ }
```

## Why it matters

Histograms summarize distributions (latencies, sizes). Integer division maps a
value to its bucket in O(1); pre-sizing avoids bounds issues.

## Watch out

- Guard `size > 0` to avoid divide-by-zero.
- Allocate `maxBin+1` bins, or the highest value indexes out of range.
- This assumes non-negative values; negatives need an offset.
