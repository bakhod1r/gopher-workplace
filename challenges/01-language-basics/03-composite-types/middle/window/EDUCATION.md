# Rolling window sums

## The idea

Naively each window sum is O(k); rolling makes it O(1) per step by adding the new
element and subtracting the one that left:

```go
sum := 0
for i := 0; i < k; i++ { sum += xs[i] }
out = append(out, sum)
for i := k; i < len(xs); i++ { sum += xs[i] - xs[i-k]; out = append(out, sum) }
```

## Why it matters

Moving averages, rate windows, and streaming stats use this. The rolling update
turns an O(n·k) computation into O(n).

## Watch out

- There are `len(xs)-k+1` windows; guard `k > len` and `k <= 0`.
- Rolling avoids recomputation but accumulates rounding for floats.
- Sub-slices would share memory; here you output computed sums.
